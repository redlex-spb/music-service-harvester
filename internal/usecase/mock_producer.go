package usecase

type mockProducer struct {
	topic string
	val   any
	err   error
}

func (m *mockProducer) Send(t string, v any) error { m.topic, m.val = t, v; return m.err }
func (m *mockProducer) Close() error               { return nil }
