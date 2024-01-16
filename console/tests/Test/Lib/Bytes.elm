module Test.Lib.Bytes exposing (..)

import Bytes.Encode as E
import Expect exposing (Expectation)
import Fuzz exposing (Fuzzer, int, list, string)
import Lib.Bytes
import Test exposing (..)


suite : Test
suite =
    describe "Test bytes decoder"
        [ test "static" <|
            \() ->
                let
                    str =
                        "abcdefghi"

                    expected =
                        List.map Char.toCode <| String.toList str
                in
                Lib.Bytes.bytesToList (E.encode (E.string str)) |> Expect.equal expected
        ]
