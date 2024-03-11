package lib

func SafeGoAndAutoRestart(f func()) {
	ch := make(chan struct{}, 1)
	fc := func() {
		defer func() {
			if err := recover(); err != nil {
				ch <- struct{}{}
			}
		}()
		f()
	}
	go func() {
		for range ch {
			go fc()
		}
	}()
	ch <- struct{}{}
}

var jobs = []func(){
	// ...
}

func StartJob() {
	for _, job := range jobs {
		SafeGoAndAutoRestart(job)
	}
}
