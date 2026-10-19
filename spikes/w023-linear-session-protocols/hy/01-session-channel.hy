package main

// W-023: Typestate Session Protocol Demonstration
// Compile-time checks ensure Ping -> Pong -> Close ordering with 0 runtime overhead.

fn client(ch chan[send int -> recv int -> close]) {
    ch <- 100
    res := <-ch
    println(res)
    close(ch)
}

fn server(ch chan[recv int -> send int]) {
    req := <-ch
    ch <- req * 2
}

fn main() {
    ch := make(chan[send int -> recv int -> close], 1)
    task server(ch)
    client(ch)
}
