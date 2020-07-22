package terrors

import (
	"fmt"
	"net/http"

	"golang.org/x/xerrors"
)

// DataflowError サーバ-クライアント間エラーハンドリング用エラー
type DataflowError struct {
	// エラーコード
	ErrorCode int
	// システムエラーメッセージ(日本語)
	ErrorMessageJP string
	// システムエラーメッセージ(英語)
	ErrorMessageEN string
	// xerrors拡張用フィールド
	err error
	// それぞれでfmt.Errorf("%w", err)を記述する必要があるためxerrors使う。
	frame xerrors.Frame
}

// New DataflowErrorを生成する
func New(errorCode int) error {
	return newError(nil, errorCode, "", "")
}

// Newf DataflowErrorをエラーメッセージ付きで生成する
func Newf(errorCode int, messageJP string, messageEN string) error {
	return newError(nil, errorCode, messageJP, messageEN)
}

// Wrap エラーをDataflowエラーでラップする
func Wrap(cause error, errorCode int) error {
	return newError(cause, errorCode, "", "")
}

// Wrapf エラーをDataflowエラーで、エラーメッセージ付きでラップする
func Wrapf(cause error, errorCode int, messageJP, messageEN string) error {
	return newError(cause, errorCode, messageJP, messageEN)
}

func newError(cause error, errorCode int, errorMessageJP, errorMessageEN string) error {
	return &DataflowError{
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            cause,
		frame:          xerrors.Caller(2),
	}
}

// Stack エラーをStackする
// スタックフレームを明示的に積んでいく必要があるためエラー出力に記録したいエラーハンドリング箇所ではStackを行う
func Stack(err error) error {
	var errorCode int
	var errorMessageJP, errorMessageEN string
	var dataflowError *DataflowError
	if ok := xerrors.As(err, &dataflowError); ok {
		errorCode = dataflowError.ErrorCode
		errorMessageJP = dataflowError.ErrorMessageJP
		errorMessageEN = dataflowError.ErrorMessageEN
	} else {
		return &DataflowError{
			ErrorCode:      http.StatusInternalServerError,
			ErrorMessageJP: "エラーのコンバート時にエラーが発生しました",
			ErrorMessageEN: "Error occured at covert to original error",
			err:            err,
			frame:          xerrors.Caller(1),
		}
	}
	return &DataflowError{
		ErrorCode:      errorCode,
		ErrorMessageJP: errorMessageJP,
		ErrorMessageEN: errorMessageEN,
		err:            err,
		frame:          xerrors.Caller(1),
	}
}

// Error エラーメッセージを取得する
func (e *DataflowError) Error() string {
	return fmt.Sprintf("messageJP=%s, messageEN=%s", e.ErrorMessageJP, e.ErrorMessageEN)
}

func (e *DataflowError) Unwrap() error {
	return e.err
}

func (e *DataflowError) Format(s fmt.State, v rune) {
	xerrors.FormatError(e, s, v)
}

func (e *DataflowError) FormatError(p xerrors.Printer) error {
	p.Print(e.ErrorMessageJP, e.ErrorMessageEN)
	e.frame.Format(p)
	return e.err
}
