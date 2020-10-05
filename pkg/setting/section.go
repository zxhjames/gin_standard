package setting

import "time"

/**
* @program: src
* @description:
* @author: 占翔昊
* @create 2020-10-05 18:34
**/

type ServerSettings struct {
	RunMode string
	HttpPort string
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}


type AppSettings struct {
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
 }

 type DatabaseSettings struct {
	 DBType string
	 Username string
	 Password string
	 Host string
	 DBName string
	 TablePrefix string
	 Charset string
	 ParseTime bool
	 MaxIdleConns int
	 MaxOpenConns int
 }

 func (s *Setting) ReadSection(k string,v interface{}) error {
 	if err := s.vp.UnmarshalKey(k,v);err!=nil {
 		return err
	}
	return nil

 }