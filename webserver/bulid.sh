#! /bin/bash
# BUlid web UI

cd D:/GO/GOProjects/src03/webserver/web
go build main.go handlers.go
cp D:/GO/GoProjects/src03/webserver/web/main.exe  D:/GO/GOProjects/bin/webserver_ui/web
cp -R D:/GO/GOProjects/src03/webserver/template D:/GO/GOProjects/bin/webserver_ui