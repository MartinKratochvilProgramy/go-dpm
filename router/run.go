package router

func (r *Router) Run() {
	r.R.Run(*r.ServerAddr)
}
