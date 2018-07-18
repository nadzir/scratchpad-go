FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/github.com/nadzir/job-analysis/app
WORKDIR /go/src/github.com/nadzir/job-analysis/app/job-crawler/cmd/main

RUN go get ./
RUN go build job-crawler.go

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi

EXPOSE 8080