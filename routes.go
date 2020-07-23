package main

func (srv *Server) routes() {
	srv.router.Get("/", srv.Index())
}
