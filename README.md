# golang-study
*Golang 공부하면서 개인적으로 기록하는 Repository*

# Architecture

## Swagger-UI

* local: http://localhost:3000/swagger/index.html

## API Path
* API Group은 기본적으로 Version을 기본 prefix로 설정하고, 이후 Path를 추가한다.
 * API Path는 다음의 규칙을 가진다. /{version}/{추상화 된 API 목적}/{path}
* API는 RESTful 목적을 지향하도록 한다.

## Directory Structure

* Base on [Golang standards](https://github.com/golang-standards/project-layout)
* Base on 'Clean Architecture' and 'DDD; Domain Driven Design'

| directory path   | description               |
|------------------|---------------------------|
| `controllers`    | HTTP Handler 의 구현체로, Request 와 Response 를 다룬다. |
| `services`       | domain 레이어와 repository 레이어를 조합하여 비즈니스 로직의 순서를 결정한다. |
| `gateways`   | 외부 시스템(e.g., DB, AWS 등)과의 모든 통신을 담당한다. |
| `domains`        | 모든 비즈니스 로직(업무 규칙)을 다룬다. |
| `assets`         | 서비스에서 사용하는 이미지, 로그 등의  기타 자산 파일                 |
| `configs`         | 서비스에서 사용하는 주요 설정 파일        |
| `docs`           | Swagger 사양 및 사용자 정의 문서 |
| `router`    | API Path, HTTP Handler 정의한다. |

## Tech Stack

### Framwork

* go 1.20.3
* Echo framework