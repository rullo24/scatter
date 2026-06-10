# Scatter

A high-throughput, concurrent file-copying engine written in Go. 

Designed to bypass standard sequential I/O bottlenecks by chunking files into uniform blocks and executing out-of-order reads and writes concurrently without global lock contention.

## Architecture

The engine eliminates global mutex locking by leveraging thread-safe POSIX positional I/O system calls (`pread`/`pwrite`).

1. **Splitter:** Measures the source file and calculates chunk offsets based on a fixed block size.
2. **Worker Pool:** Spawns a pool of goroutines (start w/ `runtime.NumCPU()` but test varying lengths) that pull chunk metadata from a shared jobs channel.
3. **Positional I/O:** Workers read from the source and write directly to the destination file concurrently using `os.ReadAt()` and `os.WriteAt()`.

## CLI Usage

```bash
scatter --src=path/to/large_file.iso --dest=path/to/copy.iso --workers=8 --blocksize=4MB
