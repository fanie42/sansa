package main

func main() {
    siteRepo := inmem.NewSiteRepo()
    instrumentRepo := inmem.NewInstrumentRepo()
    systemRepo := inmem.NewSystemRepo()

    sitePres := graphql.NewSitePres()
    instrumentPres := graphql.NewInstrumentPres()
    systemPres := graphql.NewSystemPres()

    adminService := admin.NewService(
        siteRepo,
        sitePres,
        instrumentRepo,
        instrumentPres,
        systemRepo,
        systemPres,
    )

    httpController := http.NewAdminController(
        adminService,
    )
}
