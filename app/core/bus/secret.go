package bus

import "log/slog"

type Secret string

func (Secret) String() string               { return "[REDACTED]" }
func (Secret) MarshalJSON() ([]byte, error) { return []byte(`"[REDACTED]"`), nil }
func (Secret) LogValue() slog.Value         { return slog.StringValue("[REDACTED]") }
func (s Secret) Reveal() string             { return string(s) }
