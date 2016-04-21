* go-rebound
Chad Hedstrom <chad.hedstrom@upguard.com>

** What

This is a little containerizable web application for 
performing internal resiliency testing

Forked from go-appstatus by Zakk Acreman

** TODO

	* Stop and restart an org of services
    * Randomly force kill any one docker container
    * Randomly force kill docker container on a timer
    * Monitor status of services before and after fkill
    * Write logs to a file
    * View logs from HTML front end
    * Graph per container type over time
