package configuration

type Worker struct {
	MaxQueue  int `json:"max_queue"`
	MaxWorker int `json:"max_worker"`
}

func (k *Worker) Process() {

}
