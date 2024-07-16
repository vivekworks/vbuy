package vbuy

import (
    "encoding/json"
    "io"
)

func ReadJSON[T any](r io.ReadCloser, v *T) error {
    if err := json.NewDecoder(r).Decode(v); err != nil {
        return err
    }
    return nil
}

func WriteJSON[T any](w io.Writer, v *T) error {
    if err := json.NewEncoder(w).Encode(v); err != nil {
        return err
    }
    return nil
}
