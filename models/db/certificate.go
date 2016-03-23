package db

type Certificate struct {
	InputStream interface{}
	Description string
	Open        bool
	Filename    string
	Readable    bool
	File        CertFile
}

type CertFile struct {
	Path          string
	Name          string
	Parent        string
	Absolute      bool
	CanonicalPath string
	ParentFile    *CertFile
	AbsolutePath  string
	AbsoluteFile  *CertFile
	CanonicalFile *CertFile
	Directory     bool
	File          bool
	Hidden        bool
	TotalSpace    int64
	FreeSpace     int64
	UsableSpace   int64
}
