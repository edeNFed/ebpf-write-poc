# Overriding go function arguments via eBPF - Proof of concept

This project is a small program that demonstrates how to override go function arguments via eBPF instrumentation.
The `str` argument of the `worker` function is overridden by the instrumentation.
