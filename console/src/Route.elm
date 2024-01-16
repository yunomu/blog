module Route exposing (Route(..), fromUrl, path)

import Proto.Api
import Url exposing (Url)
import Url.Builder as B
import Url.Parser as P exposing ((</>), (<?>), Parser, s)
import Url.Parser.Query as Query


type Route
    = Index
    | AuthCallback
        { code : Maybe String
        , state : Maybe String
        }
    | Files
    | InitUserForm
    | NotFound Url


authCallback : Maybe String -> Maybe String -> Route
authCallback code state =
    AuthCallback { code = code, state = state }


parser : Parser (Route -> a) a
parser =
    P.oneOf
        [ P.map Index P.top
        , P.map authCallback <| s "callback" <?> Query.string "code" <?> Query.string "state"
        , P.map Files <| s "files"
        ]


catMaybes : List (Maybe a) -> List a
catMaybes ls =
    case ls of
        [] ->
            []

        (Just a) :: xs ->
            a :: catMaybes xs

        Nothing :: xs ->
            catMaybes xs


path : Route -> String
path route =
    case route of
        AuthCallback params ->
            B.absolute [ "callback" ] <|
                catMaybes
                    [ Maybe.map (B.string "code") params.code
                    , Maybe.map (B.string "state") params.state
                    ]

        Files ->
            B.absolute [ "files" ] []

        _ ->
            B.absolute [] []


fromUrl : Url -> Route
fromUrl url =
    Maybe.withDefault (NotFound url) <| P.parse parser url
