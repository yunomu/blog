port module Main exposing (main)

import Api
import Auth
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Bytes exposing (Bytes)
import Element exposing (Element)
import Element.Lazy as Lazy
import File exposing (File)
import File.Select
import Html exposing (Html)
import Http
import Json.Decode as Decoder exposing (Decoder)
import Lib.Bytes
import Proto.Api as PB
import Route exposing (Route)
import Task
import Url exposing (Url)
import Url.Builder as UrlBuilder
import View.Files
import View.Index


port storeTokens : ( String, String, String ) -> Cmd msg


port removeTokens : () -> Cmd msg


type alias Flags =
    { idToken : Maybe String
    , accessToken : Maybe String
    , refreshToken : Maybe String
    , windowWidth : Int
    , windowHeight : Int
    , authClientId : String
    , authRedirectURL : String
    , logoutRedirectURL : String
    , idp : String
    }


type Msg
    = NOP
    | UrlRequest Browser.UrlRequest
    | UrlChanged Url
    | OnResize Int Int
    | AuthResult Msg (Result Http.Error Auth.Token)
    | RedirectToLoginForm
    | RedirectToIndex
    | AuthMsg Msg Auth.Msg
    | ApiResponse Api.Request (Result Http.Error Api.Response)
    | FileUploadRequested
    | FilesViewMsg View.Files.Msg


type alias Token =
    { idToken : String
    , accessToken : String
    , refreshToken : String
    }


type alias Model =
    { key : Nav.Key
    , route : Route
    , windowSize : ( Int, Int )
    , loginFormURL : String
    , logoutURL : String
    , logoutRedirectURL : String
    , authModel : Auth.Model
    , endpoint : String
    , filesModel : View.Files.Model
    }


init : Flags -> Url -> Nav.Key -> ( Model, Cmd Msg )
init flags url key =
    let
        loginFormURL =
            UrlBuilder.crossOrigin
                flags.idp
                [ "oauth2", "authorize" ]
                [ UrlBuilder.string "response_type" "code"
                , UrlBuilder.string "client_id" flags.authClientId
                , UrlBuilder.string "redirect_uri" flags.authRedirectURL
                ]

        endpoint =
            "/v1"

        authModel =
            Auth.init flags.authClientId flags.idp flags.authRedirectURL
    in
    ( { key = key
      , route = Route.fromUrl url
      , windowSize = ( flags.windowWidth, flags.windowHeight )
      , loginFormURL = loginFormURL
      , logoutURL = UrlBuilder.crossOrigin flags.idp [ "logout" ] []
      , logoutRedirectURL = flags.logoutRedirectURL
      , endpoint = endpoint
      , authModel = authModel
      , filesModel = View.Files.init
      }
    , Cmd.batch
        [ Nav.pushUrl key (Url.toString url)
        , apiRequest endpoint authModel Api.GetUserRequest
        ]
    )


maybe : b -> (a -> b) -> Maybe a -> b
maybe default f =
    Maybe.withDefault default << Maybe.map f


maybeCmd : Maybe a -> (a -> Cmd msg) -> Cmd msg
maybeCmd ma f =
    maybe Cmd.none f ma


apiRequest : String -> Auth.Model -> Api.Request -> Cmd Msg
apiRequest endpoint authModel =
    Api.request ApiResponse endpoint authModel


apiResponse : Model -> Api.Request -> Result Http.Error Api.Response -> ( Model, Cmd Msg )
apiResponse model request result =
    case result of
        Ok res ->
            case res of
                Api.UploadResponse r ->
                    ( { model
                        | filesModel = View.Files.init
                      }
                      -- TODO redirect
                    , Cmd.none
                    )

                _ ->
                    -- TODO
                    ( model, Cmd.none )

        Err (Http.BadStatus 401) ->
            -- TODO unauthorized
            ( model, Cmd.none )

        Err (Http.BadStatus 400) ->
            case request of
                Api.GetUserRequest ->
                    -- TODO uninitialized
                    ( model, Cmd.none )

                _ ->
                    ( model, Cmd.none )

        Err err ->
            ( model, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        UrlRequest urlRequest ->
            case urlRequest of
                Browser.Internal url ->
                    ( model, Nav.pushUrl model.key (Url.toString url) )

                Browser.External href ->
                    ( model, Nav.load href )

        UrlChanged url ->
            let
                route =
                    Route.fromUrl url
            in
            case route of
                Route.AuthCallback arg ->
                    ( model
                    , maybeCmd arg.code <|
                        \code ->
                            Auth.tokenRequest RedirectToLoginForm
                                (AuthMsg RedirectToIndex)
                                model.authModel
                                (Auth.AuthorizationCode code)
                    )

                Route.Index ->
                    ( { model | route = route }
                    , Cmd.none
                    )

                Route.Files ->
                    ( { model | route = route }
                    , apiRequest model.endpoint model.authModel (Api.ListFilesRequest Nothing)
                    )

                _ ->
                    ( { model | route = route }
                    , Cmd.none
                    )

        OnResize w h ->
            ( { model | windowSize = ( w, h ) }, Cmd.none )

        AuthResult prevMsg result ->
            case result of
                Ok token ->
                    ( model
                    , Cmd.batch
                        [ storeTokens ( token.idToken, token.accessToken, token.refreshToken )
                        , Task.perform identity <| Task.succeed prevMsg
                        ]
                    )

                Err err ->
                    ( model, Cmd.none )

        RedirectToLoginForm ->
            ( model, Nav.load model.loginFormURL )

        AuthMsg prevMsg authMsg ->
            let
                ( authModel, cmd ) =
                    Auth.update (AuthResult prevMsg) authMsg model.authModel
            in
            ( { model | authModel = authModel }
            , cmd
            )

        RedirectToIndex ->
            ( model
            , Nav.pushUrl model.key <| Route.path Route.Index
            )

        ApiResponse request result ->
            apiResponse model request result

        FileUploadRequested ->
            ( model
            , apiRequest model.endpoint
                model.authModel
                (Api.UploadRequest
                    { contentType = View.Files.mime model.filesModel
                    , blob = View.Files.bytes model.filesModel
                    }
                )
            )

        FilesViewMsg msg_ ->
            let
                ( filesModel, cmd ) =
                    View.Files.update FilesViewMsg msg_ model.filesModel
            in
            ( { model | filesModel = filesModel }, cmd )

        NOP ->
            ( model, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Events.onResize OnResize


viewFiles =
    View.Files.view FilesViewMsg FileUploadRequested


view : Model -> Browser.Document Msg
view model =
    { title = "Blog console"
    , body =
        [ Element.layout [] <|
            case model.route of
                Route.Index ->
                    Lazy.lazy View.Index.view model.loginFormURL

                Route.Files ->
                    Lazy.lazy viewFiles model.filesModel

                Route.AuthCallback _ ->
                    Element.none

                Route.InitUserForm ->
                    -- TODO
                    Element.none

                Route.NotFound url ->
                    Element.text "NotFound"
        ]
    }


main : Program Flags Model Msg
main =
    Browser.application
        { init = init
        , update = update
        , subscriptions = subscriptions
        , view = view
        , onUrlChange = UrlChanged
        , onUrlRequest = UrlRequest
        }
