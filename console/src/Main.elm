port module Main exposing (main)

import Auth
import Browser
import Browser.Events as Events
import Browser.Navigation as Nav
import Element exposing (Element)
import Element.Lazy as Lazy
import Html exposing (Html)
import Http
import Json.Decode as Decoder exposing (Decoder)
import Proto.Api
import Route exposing (Route)
import Task
import Url exposing (Url)
import Url.Builder as UrlBuilder


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

        authModel =
            Auth.init flags.authClientId flags.idp flags.authRedirectURL
    in
    ( { key = key
      , route = Route.fromUrl url
      , windowSize = ( flags.windowWidth, flags.windowHeight )
      , loginFormURL = loginFormURL
      , logoutURL = UrlBuilder.crossOrigin flags.idp [ "logout" ] []
      , logoutRedirectURL = flags.logoutRedirectURL
      , authModel = authModel
      , endpoint = "/v1"
      }
    , Cmd.batch
        [ Nav.pushUrl key (Url.toString url)
        ]
    )


maybe : b -> (a -> b) -> Maybe a -> b
maybe default f =
    Maybe.withDefault default << Maybe.map f


maybeCmd : Maybe a -> (a -> Cmd msg) -> Cmd msg
maybeCmd ma f =
    maybe Cmd.none f ma


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

        NOP ->
            ( model, Cmd.none )


subscriptions : Model -> Sub Msg
subscriptions _ =
    Events.onResize OnResize


view : Model -> Browser.Document Msg
view model =
    { title = "Blog console"
    , body =
        [ Element.layout [] <|
            Element.text "hello"
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
