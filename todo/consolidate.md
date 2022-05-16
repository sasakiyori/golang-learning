1. goroutine
2. channel
3. shared memory
4. slice
5. benchmark
   ```go
   // more accurately
   func BenchmarkFunc(b *testing.B) {
       b.StopTimer()
       //  preparation
       DoSth()
       b.StartTimer()

       for (i := 0; i < b.N; i++) {
           // benchmark
       }
   }
   ```
6. cross compiling
7. context
8. pointer