package gogit

import "path/filepath"

var (
	RepoPath         = filepath.Join(".", ".gogit")
	ObjectsPath      = filepath.Join(RepoPath, "objects")
	IndexPath        = filepath.Join(RepoPath, "index")
	HeadPath         = filepath.Join(RepoPath, "HEAD")
	RefHeadsPath     = filepath.Join(RepoPath, "refs/heads")
	RefHeadsMainPath = filepath.Join(RepoPath, "refs/heads/main")

	ROOT      = ".gogit"
	OBJECTS   = "objects"
	REF_HEADS = "refs/heads"
	HEAD      = "HEAD"
	INDEX     = "index"
)
