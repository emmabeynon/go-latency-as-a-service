package main

import (
	"net/http"
	"net/http/httptest"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Latency as a service", func() {
	Describe("GET /latency", func() {
		var (
			writer  *httptest.ResponseRecorder
			request *http.Request
		)

		Context("with no duration param", func() {
			BeforeEach(func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/latency", nil)
			})

			It("returns 200 status code", func() {
				latencyServer(writer, request)
				Expect(writer.Code).To(Equal(200))
			})

			It("returns 'OK' after 500 ms", func() {
				beforeRequest := time.Now()
				latencyServer(writer, request)
				latency := time.Since(beforeRequest) / time.Millisecond
				Expect(writer.Body.Bytes()).To(ContainSubstring("OK"))
				Expect(latency).To(BeNumerically("~", 500, 20))
			})
		})

		Context("with a valid duration param", func() {
			BeforeEach(func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/latency?duration=100ms", nil)
			})

			It("returns 200 status code", func() {
				latencyServer(writer, request)
				Expect(writer.Code).To(Equal(200))
			})

			It("returns OK after 100 ms", func() {
				beforeRequest := time.Now()
				latencyServer(writer, request)
				latency := time.Since(beforeRequest) / time.Millisecond
				Expect(writer.Body.Bytes()).To(ContainSubstring("OK"))
				Expect(latency).To(BeNumerically("~", 100, 10))
			})
		})

		Context("with an invalid duration param", func() {
			BeforeEach(func() {
				writer = httptest.NewRecorder()
				request, _ = http.NewRequest("GET", "/latency?duration=onehundred", nil)
				latencyServer(writer, request)
			})

			It("returns 400 status code", func() {
				Expect(writer.Code).To(Equal(400))
			})

			It("returns 'Error: Invalid duration parameter'", func() {
				Expect(writer.Body.Bytes()).To(ContainSubstring("Error: Invalid duration parameter"))
			})
		})
	})
})
