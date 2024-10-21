package constants

type StatusAutorizacao int

const (
    Autorizado        StatusAutorizacao = 100
    Cancelada         StatusAutorizacao = 101
    Denegada          StatusAutorizacao = 110
    DenegadoEmitente  StatusAutorizacao = 301
    Denegado205       StatusAutorizacao = 205
)
