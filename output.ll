target triple = "x86_64-pc-linux-gnu"

declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

define i32 @main() {
  %2 = fadd double 42, 0.0
  %1 = alloca double
  store double %2, double* %1
  %4 = fadd double 3.1400000000000001, 0.0
  %3 = alloca double
  store double %4, double* %3
  ret i32 0
}
