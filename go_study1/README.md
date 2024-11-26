# MyApp - HTTP 서버 예제

이 프로젝트는 기본적인 HTTP 서버를 구현한 예제입니다. 간단한 RESTful 엔드포인트와 핸들러를 통해 HTTP 요청을 처리하고 JSON 데이터를 다루는 방법을 보여줍니다.

## 파일 구조

```
/main.go               # 서버 실행 파일
/myapp/app.go          # HTTP 핸들러 정의
/myapp/app_test.go     # 테스트 코드
```

### 파일별 설명

1. **main.go**  
   HTTP 서버를 실행하며, `myapp` 패키지의 `NewHttpHandler`를 사용하여 라우팅을 설정합니다.

2. **app.go**  
   주요 HTTP 핸들러 로직이 포함되어 있습니다.

   - `/` : "Hello World" 메시지 반환
   - `/bar` : 쿼리 파라미터로 전달된 이름을 이용하여 환영 메시지 반환
   - `/foo` : POST 요청으로 전달된 JSON 데이터를 기반으로 사용자 정보를 생성하고 반환

3. **app_test.go**  
   각 핸들러의 동작을 검증하는 테스트가 포함되어 있습니다.  
   [Testify](https://github.com/stretchr/testify)를 사용하여 단위 테스트를 작성하였습니다.

---

## 주요 기능

### 1. **루트 핸들러 (`/`)**

- **HTTP 메서드**: GET
- **응답**:
  - 성공: `Hello World` 텍스트 반환

### 2. **Bar 핸들러 (`/bar`)**

- **HTTP 메서드**: GET
- **파라미터**:
  - `name` (optional): 사용자 이름
- **응답**:
  - `name` 제공 시: `Hello [name]!`
  - `name` 미제공 시: `Hello World!`

### 3. **Foo 핸들러 (`/foo`)**

- **HTTP 메서드**: POST
- **요청 바디**:  
  JSON 형식의 사용자 정보
  ```json
  {
    "first_name": "string",
    "last_name": "string",
    "email": "string"
  }
  ```
- **응답**:
  - 성공: 사용자 정보와 생성 시간 포함 JSON 반환
  - 실패: 잘못된 요청 응답 (400 Bad Request)

---

## 테스트

### 1. 테스트 구조

- `TestIndexPathHandler`  
  `/` 엔드포인트에서 "Hello World" 응답을 검증합니다.

- `TestBarPathHandler_WithoutName`  
  `/bar` 엔드포인트 호출 시 기본 메시지를 검증합니다.

- `TestBarPathHandler_WithName`  
  `/bar` 호출 시 쿼리 파라미터를 포함한 응답을 검증합니다.

- `TestFooHandler_WithOutJson`  
  `/foo`에 잘못된 요청이 전달되었을 때 400 응답을 검증합니다.

- `TestFooHandler_WithJson`  
  `/foo`에 올바른 JSON 요청이 전달되었을 때 사용자 정보를 포함한 응답을 검증합니다.

### 2. 실행 방법

Go 테스트 도구를 사용하여 모든 테스트를 실행할 수 있습니다.

```bash
go test ./myapp/...
```

---

## 기술 스택

- **언어**: Go
- **라이브러리**:
  - `net/http`: HTTP 서버 및 클라이언트 기능
  - `encoding/json`: JSON 인코딩 및 디코딩
  - [Testify](https://github.com/stretchr/testify): 단위 테스트 및 어서션 지원

---
