package zjson

// NewProductionEncoderConfig returns an opinionated EncoderConfig for
// production environments.
func NewProductionEncoderConfig() EncoderConfig {
	return EncoderConfig{
		LineEnding:     DefaultLineEnding,
		EncodeTime:     RFC3339TimeEncoder,
		EncodeDuration: SecondsDurationEncoder,
	}
}
