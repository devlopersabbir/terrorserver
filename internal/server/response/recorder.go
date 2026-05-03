package response

import "net/http"

type Recorder struct {
	http.ResponseWriter
	Code    int
	written bool
}

func NewRecorder(w http.ResponseWriter) *Recorder {
	return &Recorder{
		ResponseWriter: w,
		Code:           http.StatusOK,
	}
}

func (rw *Recorder) WriteHeader(code int) {
	if !rw.written {
		rw.Code = code
		rw.written = true
		rw.ResponseWriter.WriteHeader(code)
	}
}

func (rw *Recorder) Write(b []byte) (int, error) {
	if !rw.written {
		rw.written = true
	}
	return rw.ResponseWriter.Write(b)
}
