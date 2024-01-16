module View.Index exposing (view)

import Element exposing (Element)


view : String -> Element msg
view loginFormURL =
    Element.column []
        [ Element.link []
            { url = loginFormURL
            , label = Element.text "Login"
            }
        ]
