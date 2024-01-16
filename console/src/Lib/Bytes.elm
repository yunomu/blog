module Lib.Bytes exposing (bytesToList)

import Bytes exposing (Bytes)
import Bytes.Decode as D
import Protobuf


decodeList : Int -> D.Decoder a -> D.Decoder (List a)
decodeList len decoder =
    D.loop ( len, [] ) <|
        \( l, accm ) ->
            if l <= 0 then
                D.succeed <| D.Done <| List.reverse accm

            else
                D.map (\a -> D.Loop ( l - 1, a :: accm )) decoder


bytesToList : Bytes -> Protobuf.Bytes
bytesToList bytes =
    Maybe.withDefault [] <| D.decode (decodeList (Bytes.width bytes) D.unsignedInt8) bytes
