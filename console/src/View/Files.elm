module View.Files exposing (Model, Msg, bytes, continuationToken, init, mime, update, view)

import Bytes exposing (Bytes)
import Element exposing (Element)
import File exposing (File)
import File.Select
import Html
import Html.Attributes as Attr
import Lib.Bytes
import Proto.Api as PB
import Task
import View.Atom.Button


type Msg
    = LoadRequest
    | Selected File
    | ToBytes Bytes
    | ToUrl String


type alias Model =
    { name : String
    , mime : String
    , preview : Maybe String
    , bytes : List Int
    , files : List PB.Object
    , continuationToken : Maybe String
    }


init : List PB.Object -> Maybe String -> Model
init files ctoken =
    { name = ""
    , mime = ""
    , preview = Nothing
    , bytes = []
    , files = files
    , continuationToken = ctoken
    }


mime : Model -> String
mime model =
    model.mime


bytes : Model -> List Int
bytes model =
    model.bytes


continuationToken : Model -> Maybe String
continuationToken model =
    model.continuationToken


mimes : List String
mimes =
    [ "image/jpeg", "image/png", "image/gif" ]


update : (Msg -> msg) -> Msg -> Model -> ( Model, Cmd msg )
update toMsg msg model =
    case msg of
        LoadRequest ->
            ( model
            , File.Select.file mimes (toMsg << Selected)
            )

        Selected file ->
            let
                mime_ =
                    File.mime file
            in
            ( { model
                | name = File.name file
                , mime = mime_
              }
            , Cmd.batch
                [ if List.member mime_ mimes then
                    Task.perform (toMsg << ToUrl) <| File.toUrl file

                  else
                    Cmd.none
                , Task.perform (toMsg << ToBytes) <| File.toBytes file
                ]
            )

        ToUrl url ->
            ( { model | preview = Just url }
            , Cmd.none
            )

        ToBytes bs ->
            ( { model | bytes = Lib.Bytes.bytesToList bs }
            , Cmd.none
            )


maybe : b -> (a -> b) -> Maybe a -> b
maybe b f =
    Maybe.withDefault b << Maybe.map f


maybe_ : Maybe a -> b -> (a -> b) -> b
maybe_ ma b f =
    maybe b f ma


view :
    (Msg -> msg)
    -> msg
    -> Model
    -> Element msg
view toMsg submitMsg model =
    Element.column [ Element.paddingXY 30 20, Element.spacing 10 ]
        [ Element.text "Files"
        , Element.row []
            [ View.Atom.Button.button (toMsg LoadRequest) "File Select"
            , Element.text model.name
            ]
        , maybe_ model.preview Element.none <|
            \prev ->
                Element.column []
                    [ View.Atom.Button.updateButton submitMsg "Upload"
                    , Element.html <| Html.img [ Attr.src prev, Attr.style "width" "400px" ] []
                    ]
        ]
