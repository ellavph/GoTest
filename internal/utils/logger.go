package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"runtime"
	"time"
)

// LogLevel representa o nível de log
type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	WARN  LogLevel = "WARN"
	ERROR LogLevel = "ERROR"
	FATAL LogLevel = "FATAL"
)

// Logger representa um logger estruturado
type Logger struct {
	level  LogLevel
	output *log.Logger
}

// LogEntry representa uma entrada de log
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     LogLevel               `json:"level"`
	Message   string                 `json:"message"`
	Context   map[string]interface{} `json:"context,omitempty"`
	File      string                 `json:"file,omitempty"`
	Line      int                    `json:"line,omitempty"`
	Function  string                 `json:"function,omitempty"`
}

// NewLogger cria um novo logger
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level:  level,
		output: log.New(os.Stdout, "", 0),
	}
}

// shouldLog verifica se deve fazer log baseado no nível
func (l *Logger) shouldLog(level LogLevel) bool {
	levels := map[LogLevel]int{
		DEBUG: 0,
		INFO:  1,
		WARN:  2,
		ERROR: 3,
		FATAL: 4,
	}
	
	return levels[level] >= levels[l.level]
}

// log faz o log estruturado
func (l *Logger) log(level LogLevel, message string, ctx map[string]interface{}) {
	if !l.shouldLog(level) {
		return
	}
	
	// Obter informações do caller
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "unknown"
		line = 0
	}
	
	pc, _, _, ok := runtime.Caller(2)
	function := "unknown"
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			function = fn.Name()
		}
	}
	
	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Context:   ctx,
		File:      file,
		Line:      line,
		Function:  function,
	}
	
	jsonData, err := json.Marshal(entry)
	if err != nil {
		l.output.Printf("Error marshaling log entry: %v", err)
		return
	}
	
	l.output.Println(string(jsonData))
	
	// Se for FATAL, encerra o programa
	if level == FATAL {
		os.Exit(1)
	}
}

// Debug faz log de debug
func (l *Logger) Debug(message string, ctx ...map[string]interface{}) {
	context := make(map[string]interface{})
	if len(ctx) > 0 {
		context = ctx[0]
	}
	l.log(DEBUG, message, context)
}

// Info faz log de informação
func (l *Logger) Info(message string, ctx ...map[string]interface{}) {
	context := make(map[string]interface{})
	if len(ctx) > 0 {
		context = ctx[0]
	}
	l.log(INFO, message, context)
}

// Warn faz log de warning
func (l *Logger) Warn(message string, ctx ...map[string]interface{}) {
	context := make(map[string]interface{})
	if len(ctx) > 0 {
		context = ctx[0]
	}
	l.log(WARN, message, context)
}

// Error faz log de erro
func (l *Logger) Error(message string, err error, ctx ...map[string]interface{}) {
	context := make(map[string]interface{})
	if len(ctx) > 0 {
		context = ctx[0]
	}
	if err != nil {
		context["error"] = err.Error()
	}
	l.log(ERROR, message, context)
}

// Fatal faz log fatal e encerra o programa
func (l *Logger) Fatal(message string, err error, ctx ...map[string]interface{}) {
	context := make(map[string]interface{})
	if len(ctx) > 0 {
		context = ctx[0]
	}
	if err != nil {
		context["error"] = err.Error()
	}
	l.log(FATAL, message, context)
}

// WithContext adiciona contexto ao logger
func (l *Logger) WithContext(ctx context.Context) *ContextLogger {
	return &ContextLogger{
		logger: l,
		ctx:    ctx,
	}
}

// ContextLogger é um logger que carrega contexto
type ContextLogger struct {
	logger *Logger
	ctx    context.Context
}

// Debug faz log de debug com contexto
func (cl *ContextLogger) Debug(message string, ctx ...map[string]interface{}) {
	context := cl.extractContext()
	if len(ctx) > 0 {
		for k, v := range ctx[0] {
			context[k] = v
		}
	}
	cl.logger.log(DEBUG, message, context)
}

// Info faz log de informação com contexto
func (cl *ContextLogger) Info(message string, ctx ...map[string]interface{}) {
	context := cl.extractContext()
	if len(ctx) > 0 {
		for k, v := range ctx[0] {
			context[k] = v
		}
	}
	cl.logger.log(INFO, message, context)
}

// Error faz log de erro com contexto
func (cl *ContextLogger) Error(message string, err error, ctx ...map[string]interface{}) {
	context := cl.extractContext()
	if len(ctx) > 0 {
		for k, v := range ctx[0] {
			context[k] = v
		}
	}
	if err != nil {
		context["error"] = err.Error()
	}
	cl.logger.log(ERROR, message, context)
}

// extractContext extrai informações relevantes do contexto
func (cl *ContextLogger) extractContext() map[string]interface{} {
	context := make(map[string]interface{})
	
	// Extrair request ID se existir
	if requestID := cl.ctx.Value("request_id"); requestID != nil {
		context["request_id"] = requestID
	}
	
	// Extrair user ID se existir
	if userID := cl.ctx.Value("user_id"); userID != nil {
		context["user_id"] = userID
	}
	
	return context
}

// Instância global do logger
var defaultLogger = NewLogger(INFO)

// Funções globais para facilitar o uso
func Debug(message string, ctx ...map[string]interface{}) {
	defaultLogger.Debug(message, ctx...)
}

func Info(message string, ctx ...map[string]interface{}) {
	defaultLogger.Info(message, ctx...)
}

func Warn(message string, ctx ...map[string]interface{}) {
	defaultLogger.Warn(message, ctx...)
}

func Error(message string, err error, ctx ...map[string]interface{}) {
	defaultLogger.Error(message, err, ctx...)
}

func Fatal(message string, err error, ctx ...map[string]interface{}) {
	defaultLogger.Fatal(message, err, ctx...)
}

// SetLogLevel define o nível de log global
func SetLogLevel(level LogLevel) {
	defaultLogger.level = level
}