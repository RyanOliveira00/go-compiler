target triple = "x86_64-pc-linux-gnu"

declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

define i32 @main() {
  %2 = fadd double 42, 0.0
  %1 = alloca double
  store double %2, double* %1
  ret i32 0
}
