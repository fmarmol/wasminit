package main

var (
	FileJS = "Ly8gQ29weXJpZ2h0IDIwMTggVGhlIEdvIEF1dGhvcnMuIEFsbCByaWdodHMgcmVzZXJ2ZWQuCi8vIFVzZSBvZiB0aGlzIHNvdXJjZSBjb2RlIGlzIGdvdmVybmVkIGJ5IGEgQlNELXN0eWxlCi8vIGxpY2Vuc2UgdGhhdCBjYW4gYmUgZm91bmQgaW4gdGhlIExJQ0VOU0UgZmlsZS4KCiJ1c2Ugc3RyaWN0IjsKCigoKSA9PiB7Cgljb25zdCBlbm9zeXMgPSAoKSA9PiB7CgkJY29uc3QgZXJyID0gbmV3IEVycm9yKCJub3QgaW1wbGVtZW50ZWQiKTsKCQllcnIuY29kZSA9ICJFTk9TWVMiOwoJCXJldHVybiBlcnI7Cgl9OwoKCWlmICghZ2xvYmFsVGhpcy5mcykgewoJCWxldCBvdXRwdXRCdWYgPSAiIjsKCQlnbG9iYWxUaGlzLmZzID0gewoJCQljb25zdGFudHM6IHsgT19XUk9OTFk6IC0xLCBPX1JEV1I6IC0xLCBPX0NSRUFUOiAtMSwgT19UUlVOQzogLTEsIE9fQVBQRU5EOiAtMSwgT19FWENMOiAtMSwgT19ESVJFQ1RPUlk6IC0xIH0sIC8vIHVudXNlZAoJCQl3cml0ZVN5bmMoZmQsIGJ1ZikgewoJCQkJb3V0cHV0QnVmICs9IGRlY29kZXIuZGVjb2RlKGJ1Zik7CgkJCQljb25zdCBubCA9IG91dHB1dEJ1Zi5sYXN0SW5kZXhPZigiXG4iKTsKCQkJCWlmIChubCAhPSAtMSkgewoJCQkJCWNvbnNvbGUubG9nKG91dHB1dEJ1Zi5zdWJzdHJpbmcoMCwgbmwpKTsKCQkJCQlvdXRwdXRCdWYgPSBvdXRwdXRCdWYuc3Vic3RyaW5nKG5sICsgMSk7CgkJCQl9CgkJCQlyZXR1cm4gYnVmLmxlbmd0aDsKCQkJfSwKCQkJd3JpdGUoZmQsIGJ1Ziwgb2Zmc2V0LCBsZW5ndGgsIHBvc2l0aW9uLCBjYWxsYmFjaykgewoJCQkJaWYgKG9mZnNldCAhPT0gMCB8fCBsZW5ndGggIT09IGJ1Zi5sZW5ndGggfHwgcG9zaXRpb24gIT09IG51bGwpIHsKCQkJCQljYWxsYmFjayhlbm9zeXMoKSk7CgkJCQkJcmV0dXJuOwoJCQkJfQoJCQkJY29uc3QgbiA9IHRoaXMud3JpdGVTeW5jKGZkLCBidWYpOwoJCQkJY2FsbGJhY2sobnVsbCwgbik7CgkJCX0sCgkJCWNobW9kKHBhdGgsIG1vZGUsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJY2hvd24ocGF0aCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJY2xvc2UoZmQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJZmNobW9kKGZkLCBtb2RlLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCWZjaG93bihmZCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJZnN0YXQoZmQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJZnN5bmMoZmQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKG51bGwpOyB9LAoJCQlmdHJ1bmNhdGUoZmQsIGxlbmd0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LAoJCQlsY2hvd24ocGF0aCwgdWlkLCBnaWQsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJbGluayhwYXRoLCBsaW5rLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCWxzdGF0KHBhdGgsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJbWtkaXIocGF0aCwgcGVybSwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LAoJCQlvcGVuKHBhdGgsIGZsYWdzLCBtb2RlLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXJlYWQoZmQsIGJ1ZmZlciwgb2Zmc2V0LCBsZW5ndGgsIHBvc2l0aW9uLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXJlYWRkaXIocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LAoJCQlyZWFkbGluayhwYXRoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXJlbmFtZShmcm9tLCB0bywgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LAoJCQlybWRpcihwYXRoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXN0YXQocGF0aCwgY2FsbGJhY2spIHsgY2FsbGJhY2soZW5vc3lzKCkpOyB9LAoJCQlzeW1saW5rKHBhdGgsIGxpbmssIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQkJdHJ1bmNhdGUocGF0aCwgbGVuZ3RoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXVubGluayhwYXRoLCBjYWxsYmFjaykgeyBjYWxsYmFjayhlbm9zeXMoKSk7IH0sCgkJCXV0aW1lcyhwYXRoLCBhdGltZSwgbXRpbWUsIGNhbGxiYWNrKSB7IGNhbGxiYWNrKGVub3N5cygpKTsgfSwKCQl9OwoJfQoKCWlmICghZ2xvYmFsVGhpcy5wcm9jZXNzKSB7CgkJZ2xvYmFsVGhpcy5wcm9jZXNzID0gewoJCQlnZXR1aWQoKSB7IHJldHVybiAtMTsgfSwKCQkJZ2V0Z2lkKCkgeyByZXR1cm4gLTE7IH0sCgkJCWdldGV1aWQoKSB7IHJldHVybiAtMTsgfSwKCQkJZ2V0ZWdpZCgpIHsgcmV0dXJuIC0xOyB9LAoJCQlnZXRncm91cHMoKSB7IHRocm93IGVub3N5cygpOyB9LAoJCQlwaWQ6IC0xLAoJCQlwcGlkOiAtMSwKCQkJdW1hc2soKSB7IHRocm93IGVub3N5cygpOyB9LAoJCQljd2QoKSB7IHRocm93IGVub3N5cygpOyB9LAoJCQljaGRpcigpIHsgdGhyb3cgZW5vc3lzKCk7IH0sCgkJfQoJfQoKCWlmICghZ2xvYmFsVGhpcy5wYXRoKSB7CgkJZ2xvYmFsVGhpcy5wYXRoID0gewoJCQlyZXNvbHZlKC4uLnBhdGhTZWdtZW50cykgewoJCQkJcmV0dXJuIHBhdGhTZWdtZW50cy5qb2luKCIvIik7CgkJCX0KCQl9Cgl9CgoJaWYgKCFnbG9iYWxUaGlzLmNyeXB0bykgewoJCXRocm93IG5ldyBFcnJvcigiZ2xvYmFsVGhpcy5jcnlwdG8gaXMgbm90IGF2YWlsYWJsZSwgcG9seWZpbGwgcmVxdWlyZWQgKGNyeXB0by5nZXRSYW5kb21WYWx1ZXMgb25seSkiKTsKCX0KCglpZiAoIWdsb2JhbFRoaXMucGVyZm9ybWFuY2UpIHsKCQl0aHJvdyBuZXcgRXJyb3IoImdsb2JhbFRoaXMucGVyZm9ybWFuY2UgaXMgbm90IGF2YWlsYWJsZSwgcG9seWZpbGwgcmVxdWlyZWQgKHBlcmZvcm1hbmNlLm5vdyBvbmx5KSIpOwoJfQoKCWlmICghZ2xvYmFsVGhpcy5UZXh0RW5jb2RlcikgewoJCXRocm93IG5ldyBFcnJvcigiZ2xvYmFsVGhpcy5UZXh0RW5jb2RlciBpcyBub3QgYXZhaWxhYmxlLCBwb2x5ZmlsbCByZXF1aXJlZCIpOwoJfQoKCWlmICghZ2xvYmFsVGhpcy5UZXh0RGVjb2RlcikgewoJCXRocm93IG5ldyBFcnJvcigiZ2xvYmFsVGhpcy5UZXh0RGVjb2RlciBpcyBub3QgYXZhaWxhYmxlLCBwb2x5ZmlsbCByZXF1aXJlZCIpOwoJfQoKCWNvbnN0IGVuY29kZXIgPSBuZXcgVGV4dEVuY29kZXIoInV0Zi04Iik7Cgljb25zdCBkZWNvZGVyID0gbmV3IFRleHREZWNvZGVyKCJ1dGYtOCIpOwoKCWdsb2JhbFRoaXMuR28gPSBjbGFzcyB7CgkJY29uc3RydWN0b3IoKSB7CgkJCXRoaXMuYXJndiA9IFsianMiXTsKCQkJdGhpcy5lbnYgPSB7fTsKCQkJdGhpcy5leGl0ID0gKGNvZGUpID0+IHsKCQkJCWlmIChjb2RlICE9PSAwKSB7CgkJCQkJY29uc29sZS53YXJuKCJleGl0IGNvZGU6IiwgY29kZSk7CgkJCQl9CgkJCX07CgkJCXRoaXMuX2V4aXRQcm9taXNlID0gbmV3IFByb21pc2UoKHJlc29sdmUpID0+IHsKCQkJCXRoaXMuX3Jlc29sdmVFeGl0UHJvbWlzZSA9IHJlc29sdmU7CgkJCX0pOwoJCQl0aGlzLl9wZW5kaW5nRXZlbnQgPSBudWxsOwoJCQl0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cyA9IG5ldyBNYXAoKTsKCQkJdGhpcy5fbmV4dENhbGxiYWNrVGltZW91dElEID0gMTsKCgkJCWNvbnN0IHNldEludDY0ID0gKGFkZHIsIHYpID0+IHsKCQkJCXRoaXMubWVtLnNldFVpbnQzMihhZGRyICsgMCwgdiwgdHJ1ZSk7CgkJCQl0aGlzLm1lbS5zZXRVaW50MzIoYWRkciArIDQsIE1hdGguZmxvb3IodiAvIDQyOTQ5NjcyOTYpLCB0cnVlKTsKCQkJfQoKCQkJY29uc3Qgc2V0SW50MzIgPSAoYWRkciwgdikgPT4gewoJCQkJdGhpcy5tZW0uc2V0VWludDMyKGFkZHIgKyAwLCB2LCB0cnVlKTsKCQkJfQoKCQkJY29uc3QgZ2V0SW50NjQgPSAoYWRkcikgPT4gewoJCQkJY29uc3QgbG93ID0gdGhpcy5tZW0uZ2V0VWludDMyKGFkZHIgKyAwLCB0cnVlKTsKCQkJCWNvbnN0IGhpZ2ggPSB0aGlzLm1lbS5nZXRJbnQzMihhZGRyICsgNCwgdHJ1ZSk7CgkJCQlyZXR1cm4gbG93ICsgaGlnaCAqIDQyOTQ5NjcyOTY7CgkJCX0KCgkJCWNvbnN0IGxvYWRWYWx1ZSA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBmID0gdGhpcy5tZW0uZ2V0RmxvYXQ2NChhZGRyLCB0cnVlKTsKCQkJCWlmIChmID09PSAwKSB7CgkJCQkJcmV0dXJuIHVuZGVmaW5lZDsKCQkJCX0KCQkJCWlmICghaXNOYU4oZikpIHsKCQkJCQlyZXR1cm4gZjsKCQkJCX0KCgkJCQljb25zdCBpZCA9IHRoaXMubWVtLmdldFVpbnQzMihhZGRyLCB0cnVlKTsKCQkJCXJldHVybiB0aGlzLl92YWx1ZXNbaWRdOwoJCQl9CgoJCQljb25zdCBzdG9yZVZhbHVlID0gKGFkZHIsIHYpID0+IHsKCQkJCWNvbnN0IG5hbkhlYWQgPSAweDdGRjgwMDAwOwoKCQkJCWlmICh0eXBlb2YgdiA9PT0gIm51bWJlciIgJiYgdiAhPT0gMCkgewoJCQkJCWlmIChpc05hTih2KSkgewoJCQkJCQl0aGlzLm1lbS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQsIHRydWUpOwoJCQkJCQl0aGlzLm1lbS5zZXRVaW50MzIoYWRkciwgMCwgdHJ1ZSk7CgkJCQkJCXJldHVybjsKCQkJCQl9CgkJCQkJdGhpcy5tZW0uc2V0RmxvYXQ2NChhZGRyLCB2LCB0cnVlKTsKCQkJCQlyZXR1cm47CgkJCQl9CgoJCQkJaWYgKHYgPT09IHVuZGVmaW5lZCkgewoJCQkJCXRoaXMubWVtLnNldEZsb2F0NjQoYWRkciwgMCwgdHJ1ZSk7CgkJCQkJcmV0dXJuOwoJCQkJfQoKCQkJCWxldCBpZCA9IHRoaXMuX2lkcy5nZXQodik7CgkJCQlpZiAoaWQgPT09IHVuZGVmaW5lZCkgewoJCQkJCWlkID0gdGhpcy5faWRQb29sLnBvcCgpOwoJCQkJCWlmIChpZCA9PT0gdW5kZWZpbmVkKSB7CgkJCQkJCWlkID0gdGhpcy5fdmFsdWVzLmxlbmd0aDsKCQkJCQl9CgkJCQkJdGhpcy5fdmFsdWVzW2lkXSA9IHY7CgkJCQkJdGhpcy5fZ29SZWZDb3VudHNbaWRdID0gMDsKCQkJCQl0aGlzLl9pZHMuc2V0KHYsIGlkKTsKCQkJCX0KCQkJCXRoaXMuX2dvUmVmQ291bnRzW2lkXSsrOwoJCQkJbGV0IHR5cGVGbGFnID0gMDsKCQkJCXN3aXRjaCAodHlwZW9mIHYpIHsKCQkJCQljYXNlICJvYmplY3QiOgoJCQkJCQlpZiAodiAhPT0gbnVsbCkgewoJCQkJCQkJdHlwZUZsYWcgPSAxOwoJCQkJCQl9CgkJCQkJCWJyZWFrOwoJCQkJCWNhc2UgInN0cmluZyI6CgkJCQkJCXR5cGVGbGFnID0gMjsKCQkJCQkJYnJlYWs7CgkJCQkJY2FzZSAic3ltYm9sIjoKCQkJCQkJdHlwZUZsYWcgPSAzOwoJCQkJCQlicmVhazsKCQkJCQljYXNlICJmdW5jdGlvbiI6CgkJCQkJCXR5cGVGbGFnID0gNDsKCQkJCQkJYnJlYWs7CgkJCQl9CgkJCQl0aGlzLm1lbS5zZXRVaW50MzIoYWRkciArIDQsIG5hbkhlYWQgfCB0eXBlRmxhZywgdHJ1ZSk7CgkJCQl0aGlzLm1lbS5zZXRVaW50MzIoYWRkciwgaWQsIHRydWUpOwoJCQl9CgoJCQljb25zdCBsb2FkU2xpY2UgPSAoYWRkcikgPT4gewoJCQkJY29uc3QgYXJyYXkgPSBnZXRJbnQ2NChhZGRyICsgMCk7CgkJCQljb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7CgkJCQlyZXR1cm4gbmV3IFVpbnQ4QXJyYXkodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIGFycmF5LCBsZW4pOwoJCQl9CgoJCQljb25zdCBsb2FkU2xpY2VPZlZhbHVlcyA9IChhZGRyKSA9PiB7CgkJCQljb25zdCBhcnJheSA9IGdldEludDY0KGFkZHIgKyAwKTsKCQkJCWNvbnN0IGxlbiA9IGdldEludDY0KGFkZHIgKyA4KTsKCQkJCWNvbnN0IGEgPSBuZXcgQXJyYXkobGVuKTsKCQkJCWZvciAobGV0IGkgPSAwOyBpIDwgbGVuOyBpKyspIHsKCQkJCQlhW2ldID0gbG9hZFZhbHVlKGFycmF5ICsgaSAqIDgpOwoJCQkJfQoJCQkJcmV0dXJuIGE7CgkJCX0KCgkJCWNvbnN0IGxvYWRTdHJpbmcgPSAoYWRkcikgPT4gewoJCQkJY29uc3Qgc2FkZHIgPSBnZXRJbnQ2NChhZGRyICsgMCk7CgkJCQljb25zdCBsZW4gPSBnZXRJbnQ2NChhZGRyICsgOCk7CgkJCQlyZXR1cm4gZGVjb2Rlci5kZWNvZGUobmV3IERhdGFWaWV3KHRoaXMuX2luc3QuZXhwb3J0cy5tZW0uYnVmZmVyLCBzYWRkciwgbGVuKSk7CgkJCX0KCgkJCWNvbnN0IHRlc3RDYWxsRXhwb3J0ID0gKGEsIGIpID0+IHsKCQkJCXRoaXMuX2luc3QuZXhwb3J0cy50ZXN0RXhwb3J0MCgpOwoJCQkJcmV0dXJuIHRoaXMuX2luc3QuZXhwb3J0cy50ZXN0RXhwb3J0KGEsIGIpOwoJCQl9CgoJCQljb25zdCB0aW1lT3JpZ2luID0gRGF0ZS5ub3coKSAtIHBlcmZvcm1hbmNlLm5vdygpOwoJCQl0aGlzLmltcG9ydE9iamVjdCA9IHsKCQkJCV9nb3Rlc3Q6IHsKCQkJCQlhZGQ6IChhLCBiKSA9PiBhICsgYiwKCQkJCQljYWxsRXhwb3J0OiB0ZXN0Q2FsbEV4cG9ydCwKCQkJCX0sCgkJCQlnb2pzOiB7CgkJCQkJLy8gR28ncyBTUCBkb2VzIG5vdCBjaGFuZ2UgYXMgbG9uZyBhcyBubyBHbyBjb2RlIGlzIHJ1bm5pbmcuIFNvbWUgb3BlcmF0aW9ucyAoZS5nLiBjYWxscywgZ2V0dGVycyBhbmQgc2V0dGVycykKCQkJCQkvLyBtYXkgc3luY2hyb25vdXNseSB0cmlnZ2VyIGEgR28gZXZlbnQgaGFuZGxlci4gVGhpcyBtYWtlcyBHbyBjb2RlIGdldCBleGVjdXRlZCBpbiB0aGUgbWlkZGxlIG9mIHRoZSBpbXBvcnRlZAoJCQkJCS8vIGZ1bmN0aW9uLiBBIGdvcm91dGluZSBjYW4gc3dpdGNoIHRvIGEgbmV3IHN0YWNrIGlmIHRoZSBjdXJyZW50IHN0YWNrIGlzIHRvbyBzbWFsbCAoc2VlIG1vcmVzdGFjayBmdW5jdGlvbikuCgkJCQkJLy8gVGhpcyBjaGFuZ2VzIHRoZSBTUCwgdGh1cyB3ZSBoYXZlIHRvIHVwZGF0ZSB0aGUgU1AgdXNlZCBieSB0aGUgaW1wb3J0ZWQgZnVuY3Rpb24uCgoJCQkJCS8vIGZ1bmMgd2FzbUV4aXQoY29kZSBpbnQzMikKCQkJCQkicnVudGltZS53YXNtRXhpdCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IGNvZGUgPSB0aGlzLm1lbS5nZXRJbnQzMihzcCArIDgsIHRydWUpOwoJCQkJCQl0aGlzLmV4aXRlZCA9IHRydWU7CgkJCQkJCWRlbGV0ZSB0aGlzLl9pbnN0OwoJCQkJCQlkZWxldGUgdGhpcy5fdmFsdWVzOwoJCQkJCQlkZWxldGUgdGhpcy5fZ29SZWZDb3VudHM7CgkJCQkJCWRlbGV0ZSB0aGlzLl9pZHM7CgkJCQkJCWRlbGV0ZSB0aGlzLl9pZFBvb2w7CgkJCQkJCXRoaXMuZXhpdChjb2RlKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHdhc21Xcml0ZShmZCB1aW50cHRyLCBwIHVuc2FmZS5Qb2ludGVyLCBuIGludDMyKQoJCQkJCSJydW50aW1lLndhc21Xcml0ZSI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IGZkID0gZ2V0SW50NjQoc3AgKyA4KTsKCQkJCQkJY29uc3QgcCA9IGdldEludDY0KHNwICsgMTYpOwoJCQkJCQljb25zdCBuID0gdGhpcy5tZW0uZ2V0SW50MzIoc3AgKyAyNCwgdHJ1ZSk7CgkJCQkJCWZzLndyaXRlU3luYyhmZCwgbmV3IFVpbnQ4QXJyYXkodGhpcy5faW5zdC5leHBvcnRzLm1lbS5idWZmZXIsIHAsIG4pKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHJlc2V0TWVtb3J5RGF0YVZpZXcoKQoJCQkJCSJydW50aW1lLnJlc2V0TWVtb3J5RGF0YVZpZXciOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQl0aGlzLm1lbSA9IG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlcik7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBuYW5vdGltZTEoKSBpbnQ2NAoJCQkJCSJydW50aW1lLm5hbm90aW1lMSI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCXNldEludDY0KHNwICsgOCwgKHRpbWVPcmlnaW4gKyBwZXJmb3JtYW5jZS5ub3coKSkgKiAxMDAwMDAwKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHdhbGx0aW1lKCkgKHNlYyBpbnQ2NCwgbnNlYyBpbnQzMikKCQkJCQkicnVudGltZS53YWxsdGltZSI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IG1zZWMgPSAobmV3IERhdGUpLmdldFRpbWUoKTsKCQkJCQkJc2V0SW50NjQoc3AgKyA4LCBtc2VjIC8gMTAwMCk7CgkJCQkJCXRoaXMubWVtLnNldEludDMyKHNwICsgMTYsIChtc2VjICUgMTAwMCkgKiAxMDAwMDAwLCB0cnVlKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIHNjaGVkdWxlVGltZW91dEV2ZW50KGRlbGF5IGludDY0KSBpbnQzMgoJCQkJCSJydW50aW1lLnNjaGVkdWxlVGltZW91dEV2ZW50IjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJY29uc3QgaWQgPSB0aGlzLl9uZXh0Q2FsbGJhY2tUaW1lb3V0SUQ7CgkJCQkJCXRoaXMuX25leHRDYWxsYmFja1RpbWVvdXRJRCsrOwoJCQkJCQl0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cy5zZXQoaWQsIHNldFRpbWVvdXQoCgkJCQkJCQkoKSA9PiB7CgkJCQkJCQkJdGhpcy5fcmVzdW1lKCk7CgkJCQkJCQkJd2hpbGUgKHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLmhhcyhpZCkpIHsKCQkJCQkJCQkJLy8gZm9yIHNvbWUgcmVhc29uIEdvIGZhaWxlZCB0byByZWdpc3RlciB0aGUgdGltZW91dCBldmVudCwgbG9nIGFuZCB0cnkgYWdhaW4KCQkJCQkJCQkJLy8gKHRlbXBvcmFyeSB3b3JrYXJvdW5kIGZvciBodHRwczovL2dpdGh1Yi5jb20vZ29sYW5nL2dvL2lzc3Vlcy8yODk3NSkKCQkJCQkJCQkJY29uc29sZS53YXJuKCJzY2hlZHVsZVRpbWVvdXRFdmVudDogbWlzc2VkIHRpbWVvdXQgZXZlbnQiKTsKCQkJCQkJCQkJdGhpcy5fcmVzdW1lKCk7CgkJCQkJCQkJfQoJCQkJCQkJfSwKCQkJCQkJCWdldEludDY0KHNwICsgOCksCgkJCQkJCSkpOwoJCQkJCQl0aGlzLm1lbS5zZXRJbnQzMihzcCArIDE2LCBpZCwgdHJ1ZSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBjbGVhclRpbWVvdXRFdmVudChpZCBpbnQzMikKCQkJCQkicnVudGltZS5jbGVhclRpbWVvdXRFdmVudCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IGlkID0gdGhpcy5tZW0uZ2V0SW50MzIoc3AgKyA4LCB0cnVlKTsKCQkJCQkJY2xlYXJUaW1lb3V0KHRoaXMuX3NjaGVkdWxlZFRpbWVvdXRzLmdldChpZCkpOwoJCQkJCQl0aGlzLl9zY2hlZHVsZWRUaW1lb3V0cy5kZWxldGUoaWQpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgZ2V0UmFuZG9tRGF0YShyIFtdYnl0ZSkKCQkJCQkicnVudGltZS5nZXRSYW5kb21EYXRhIjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJY3J5cHRvLmdldFJhbmRvbVZhbHVlcyhsb2FkU2xpY2Uoc3AgKyA4KSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBmaW5hbGl6ZVJlZih2IHJlZikKCQkJCQkic3lzY2FsbC9qcy5maW5hbGl6ZVJlZiI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IGlkID0gdGhpcy5tZW0uZ2V0VWludDMyKHNwICsgOCwgdHJ1ZSk7CgkJCQkJCXRoaXMuX2dvUmVmQ291bnRzW2lkXS0tOwoJCQkJCQlpZiAodGhpcy5fZ29SZWZDb3VudHNbaWRdID09PSAwKSB7CgkJCQkJCQljb25zdCB2ID0gdGhpcy5fdmFsdWVzW2lkXTsKCQkJCQkJCXRoaXMuX3ZhbHVlc1tpZF0gPSBudWxsOwoJCQkJCQkJdGhpcy5faWRzLmRlbGV0ZSh2KTsKCQkJCQkJCXRoaXMuX2lkUG9vbC5wdXNoKGlkKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgc3RyaW5nVmFsKHZhbHVlIHN0cmluZykgcmVmCgkJCQkJInN5c2NhbGwvanMuc3RyaW5nVmFsIjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJc3RvcmVWYWx1ZShzcCArIDI0LCBsb2FkU3RyaW5nKHNwICsgOCkpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVHZXQodiByZWYsIHAgc3RyaW5nKSByZWYKCQkJCQkic3lzY2FsbC9qcy52YWx1ZUdldCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IHJlc3VsdCA9IFJlZmxlY3QuZ2V0KGxvYWRWYWx1ZShzcCArIDgpLCBsb2FkU3RyaW5nKHNwICsgMTYpKTsKCQkJCQkJc3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmUKCQkJCQkJc3RvcmVWYWx1ZShzcCArIDMyLCByZXN1bHQpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVTZXQodiByZWYsIHAgc3RyaW5nLCB4IHJlZikKCQkJCQkic3lzY2FsbC9qcy52YWx1ZVNldCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCVJlZmxlY3Quc2V0KGxvYWRWYWx1ZShzcCArIDgpLCBsb2FkU3RyaW5nKHNwICsgMTYpLCBsb2FkVmFsdWUoc3AgKyAzMikpOwoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVEZWxldGUodiByZWYsIHAgc3RyaW5nKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlRGVsZXRlIjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJUmVmbGVjdC5kZWxldGVQcm9wZXJ0eShsb2FkVmFsdWUoc3AgKyA4KSwgbG9hZFN0cmluZyhzcCArIDE2KSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUluZGV4KHYgcmVmLCBpIGludCkgcmVmCgkJCQkJInN5c2NhbGwvanMudmFsdWVJbmRleCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCXN0b3JlVmFsdWUoc3AgKyAyNCwgUmVmbGVjdC5nZXQobG9hZFZhbHVlKHNwICsgOCksIGdldEludDY0KHNwICsgMTYpKSk7CgkJCQkJfSwKCgkJCQkJLy8gdmFsdWVTZXRJbmRleCh2IHJlZiwgaSBpbnQsIHggcmVmKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlU2V0SW5kZXgiOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQlSZWZsZWN0LnNldChsb2FkVmFsdWUoc3AgKyA4KSwgZ2V0SW50NjQoc3AgKyAxNiksIGxvYWRWYWx1ZShzcCArIDI0KSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUNhbGwodiByZWYsIG0gc3RyaW5nLCBhcmdzIFtdcmVmKSAocmVmLCBib29sKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlQ2FsbCI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCXRyeSB7CgkJCQkJCQljb25zdCB2ID0gbG9hZFZhbHVlKHNwICsgOCk7CgkJCQkJCQljb25zdCBtID0gUmVmbGVjdC5nZXQodiwgbG9hZFN0cmluZyhzcCArIDE2KSk7CgkJCQkJCQljb25zdCBhcmdzID0gbG9hZFNsaWNlT2ZWYWx1ZXMoc3AgKyAzMik7CgkJCQkJCQljb25zdCByZXN1bHQgPSBSZWZsZWN0LmFwcGx5KG0sIHYsIGFyZ3MpOwoJCQkJCQkJc3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmUKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA1NiwgcmVzdWx0KTsKCQkJCQkJCXRoaXMubWVtLnNldFVpbnQ4KHNwICsgNjQsIDEpOwoJCQkJCQl9IGNhdGNoIChlcnIpIHsKCQkJCQkJCXNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlCgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNTYsIGVycik7CgkJCQkJCQl0aGlzLm1lbS5zZXRVaW50OChzcCArIDY0LCAwKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVJbnZva2UodiByZWYsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMudmFsdWVJbnZva2UiOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQl0cnkgewoJCQkJCQkJY29uc3QgdiA9IGxvYWRWYWx1ZShzcCArIDgpOwoJCQkJCQkJY29uc3QgYXJncyA9IGxvYWRTbGljZU9mVmFsdWVzKHNwICsgMTYpOwoJCQkJCQkJY29uc3QgcmVzdWx0ID0gUmVmbGVjdC5hcHBseSh2LCB1bmRlZmluZWQsIGFyZ3MpOwoJCQkJCQkJc3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmUKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA0MCwgcmVzdWx0KTsKCQkJCQkJCXRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDEpOwoJCQkJCQl9IGNhdGNoIChlcnIpIHsKCQkJCQkJCXNwID0gdGhpcy5faW5zdC5leHBvcnRzLmdldHNwKCkgPj4+IDA7IC8vIHNlZSBjb21tZW50IGFib3ZlCgkJCQkJCQlzdG9yZVZhbHVlKHNwICsgNDAsIGVycik7CgkJCQkJCQl0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTsKCQkJCQkJfQoJCQkJCX0sCgoJCQkJCS8vIGZ1bmMgdmFsdWVOZXcodiByZWYsIGFyZ3MgW11yZWYpIChyZWYsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMudmFsdWVOZXciOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQl0cnkgewoJCQkJCQkJY29uc3QgdiA9IGxvYWRWYWx1ZShzcCArIDgpOwoJCQkJCQkJY29uc3QgYXJncyA9IGxvYWRTbGljZU9mVmFsdWVzKHNwICsgMTYpOwoJCQkJCQkJY29uc3QgcmVzdWx0ID0gUmVmbGVjdC5jb25zdHJ1Y3QodiwgYXJncyk7CgkJCQkJCQlzcCA9IHRoaXMuX2luc3QuZXhwb3J0cy5nZXRzcCgpID4+PiAwOyAvLyBzZWUgY29tbWVudCBhYm92ZQoJCQkJCQkJc3RvcmVWYWx1ZShzcCArIDQwLCByZXN1bHQpOwoJCQkJCQkJdGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMSk7CgkJCQkJCX0gY2F0Y2ggKGVycikgewoJCQkJCQkJc3AgPSB0aGlzLl9pbnN0LmV4cG9ydHMuZ2V0c3AoKSA+Pj4gMDsgLy8gc2VlIGNvbW1lbnQgYWJvdmUKCQkJCQkJCXN0b3JlVmFsdWUoc3AgKyA0MCwgZXJyKTsKCQkJCQkJCXRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDApOwoJCQkJCQl9CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUxlbmd0aCh2IHJlZikgaW50CgkJCQkJInN5c2NhbGwvanMudmFsdWVMZW5ndGgiOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQlzZXRJbnQ2NChzcCArIDE2LCBwYXJzZUludChsb2FkVmFsdWUoc3AgKyA4KS5sZW5ndGgpKTsKCQkJCQl9LAoKCQkJCQkvLyB2YWx1ZVByZXBhcmVTdHJpbmcodiByZWYpIChyZWYsIGludCkKCQkJCQkic3lzY2FsbC9qcy52YWx1ZVByZXBhcmVTdHJpbmciOiAoc3ApID0+IHsKCQkJCQkJc3AgPj4+PSAwOwoJCQkJCQljb25zdCBzdHIgPSBlbmNvZGVyLmVuY29kZShTdHJpbmcobG9hZFZhbHVlKHNwICsgOCkpKTsKCQkJCQkJc3RvcmVWYWx1ZShzcCArIDE2LCBzdHIpOwoJCQkJCQlzZXRJbnQ2NChzcCArIDI0LCBzdHIubGVuZ3RoKTsKCQkJCQl9LAoKCQkJCQkvLyB2YWx1ZUxvYWRTdHJpbmcodiByZWYsIGIgW11ieXRlKQoJCQkJCSJzeXNjYWxsL2pzLnZhbHVlTG9hZFN0cmluZyI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IHN0ciA9IGxvYWRWYWx1ZShzcCArIDgpOwoJCQkJCQlsb2FkU2xpY2Uoc3AgKyAxNikuc2V0KHN0cik7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyB2YWx1ZUluc3RhbmNlT2YodiByZWYsIHQgcmVmKSBib29sCgkJCQkJInN5c2NhbGwvanMudmFsdWVJbnN0YW5jZU9mIjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJdGhpcy5tZW0uc2V0VWludDgoc3AgKyAyNCwgKGxvYWRWYWx1ZShzcCArIDgpIGluc3RhbmNlb2YgbG9hZFZhbHVlKHNwICsgMTYpKSA/IDEgOiAwKTsKCQkJCQl9LAoKCQkJCQkvLyBmdW5jIGNvcHlCeXRlc1RvR28oZHN0IFtdYnl0ZSwgc3JjIHJlZikgKGludCwgYm9vbCkKCQkJCQkic3lzY2FsbC9qcy5jb3B5Qnl0ZXNUb0dvIjogKHNwKSA9PiB7CgkJCQkJCXNwID4+Pj0gMDsKCQkJCQkJY29uc3QgZHN0ID0gbG9hZFNsaWNlKHNwICsgOCk7CgkJCQkJCWNvbnN0IHNyYyA9IGxvYWRWYWx1ZShzcCArIDMyKTsKCQkJCQkJaWYgKCEoc3JjIGluc3RhbmNlb2YgVWludDhBcnJheSB8fCBzcmMgaW5zdGFuY2VvZiBVaW50OENsYW1wZWRBcnJheSkpIHsKCQkJCQkJCXRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDApOwoJCQkJCQkJcmV0dXJuOwoJCQkJCQl9CgkJCQkJCWNvbnN0IHRvQ29weSA9IHNyYy5zdWJhcnJheSgwLCBkc3QubGVuZ3RoKTsKCQkJCQkJZHN0LnNldCh0b0NvcHkpOwoJCQkJCQlzZXRJbnQ2NChzcCArIDQwLCB0b0NvcHkubGVuZ3RoKTsKCQkJCQkJdGhpcy5tZW0uc2V0VWludDgoc3AgKyA0OCwgMSk7CgkJCQkJfSwKCgkJCQkJLy8gZnVuYyBjb3B5Qnl0ZXNUb0pTKGRzdCByZWYsIHNyYyBbXWJ5dGUpIChpbnQsIGJvb2wpCgkJCQkJInN5c2NhbGwvanMuY29weUJ5dGVzVG9KUyI6IChzcCkgPT4gewoJCQkJCQlzcCA+Pj49IDA7CgkJCQkJCWNvbnN0IGRzdCA9IGxvYWRWYWx1ZShzcCArIDgpOwoJCQkJCQljb25zdCBzcmMgPSBsb2FkU2xpY2Uoc3AgKyAxNik7CgkJCQkJCWlmICghKGRzdCBpbnN0YW5jZW9mIFVpbnQ4QXJyYXkgfHwgZHN0IGluc3RhbmNlb2YgVWludDhDbGFtcGVkQXJyYXkpKSB7CgkJCQkJCQl0aGlzLm1lbS5zZXRVaW50OChzcCArIDQ4LCAwKTsKCQkJCQkJCXJldHVybjsKCQkJCQkJfQoJCQkJCQljb25zdCB0b0NvcHkgPSBzcmMuc3ViYXJyYXkoMCwgZHN0Lmxlbmd0aCk7CgkJCQkJCWRzdC5zZXQodG9Db3B5KTsKCQkJCQkJc2V0SW50NjQoc3AgKyA0MCwgdG9Db3B5Lmxlbmd0aCk7CgkJCQkJCXRoaXMubWVtLnNldFVpbnQ4KHNwICsgNDgsIDEpOwoJCQkJCX0sCgoJCQkJCSJkZWJ1ZyI6ICh2YWx1ZSkgPT4gewoJCQkJCQljb25zb2xlLmxvZyh2YWx1ZSk7CgkJCQkJfSwKCQkJCX0KCQkJfTsKCQl9CgoJCWFzeW5jIHJ1bihpbnN0YW5jZSkgewoJCQlpZiAoIShpbnN0YW5jZSBpbnN0YW5jZW9mIFdlYkFzc2VtYmx5Lkluc3RhbmNlKSkgewoJCQkJdGhyb3cgbmV3IEVycm9yKCJHby5ydW46IFdlYkFzc2VtYmx5Lkluc3RhbmNlIGV4cGVjdGVkIik7CgkJCX0KCQkJdGhpcy5faW5zdCA9IGluc3RhbmNlOwoJCQl0aGlzLm1lbSA9IG5ldyBEYXRhVmlldyh0aGlzLl9pbnN0LmV4cG9ydHMubWVtLmJ1ZmZlcik7CgkJCXRoaXMuX3ZhbHVlcyA9IFsgLy8gSlMgdmFsdWVzIHRoYXQgR28gY3VycmVudGx5IGhhcyByZWZlcmVuY2VzIHRvLCBpbmRleGVkIGJ5IHJlZmVyZW5jZSBpZAoJCQkJTmFOLAoJCQkJMCwKCQkJCW51bGwsCgkJCQl0cnVlLAoJCQkJZmFsc2UsCgkJCQlnbG9iYWxUaGlzLAoJCQkJdGhpcywKCQkJXTsKCQkJdGhpcy5fZ29SZWZDb3VudHMgPSBuZXcgQXJyYXkodGhpcy5fdmFsdWVzLmxlbmd0aCkuZmlsbChJbmZpbml0eSk7IC8vIG51bWJlciBvZiByZWZlcmVuY2VzIHRoYXQgR28gaGFzIHRvIGEgSlMgdmFsdWUsIGluZGV4ZWQgYnkgcmVmZXJlbmNlIGlkCgkJCXRoaXMuX2lkcyA9IG5ldyBNYXAoWyAvLyBtYXBwaW5nIGZyb20gSlMgdmFsdWVzIHRvIHJlZmVyZW5jZSBpZHMKCQkJCVswLCAxXSwKCQkJCVtudWxsLCAyXSwKCQkJCVt0cnVlLCAzXSwKCQkJCVtmYWxzZSwgNF0sCgkJCQlbZ2xvYmFsVGhpcywgNV0sCgkJCQlbdGhpcywgNl0sCgkJCV0pOwoJCQl0aGlzLl9pZFBvb2wgPSBbXTsgICAvLyB1bnVzZWQgaWRzIHRoYXQgaGF2ZSBiZWVuIGdhcmJhZ2UgY29sbGVjdGVkCgkJCXRoaXMuZXhpdGVkID0gZmFsc2U7IC8vIHdoZXRoZXIgdGhlIEdvIHByb2dyYW0gaGFzIGV4aXRlZAoKCQkJLy8gUGFzcyBjb21tYW5kIGxpbmUgYXJndW1lbnRzIGFuZCBlbnZpcm9ubWVudCB2YXJpYWJsZXMgdG8gV2ViQXNzZW1ibHkgYnkgd3JpdGluZyB0aGVtIHRvIHRoZSBsaW5lYXIgbWVtb3J5LgoJCQlsZXQgb2Zmc2V0ID0gNDA5NjsKCgkJCWNvbnN0IHN0clB0ciA9IChzdHIpID0+IHsKCQkJCWNvbnN0IHB0ciA9IG9mZnNldDsKCQkJCWNvbnN0IGJ5dGVzID0gZW5jb2Rlci5lbmNvZGUoc3RyICsgIlwwIik7CgkJCQluZXcgVWludDhBcnJheSh0aGlzLm1lbS5idWZmZXIsIG9mZnNldCwgYnl0ZXMubGVuZ3RoKS5zZXQoYnl0ZXMpOwoJCQkJb2Zmc2V0ICs9IGJ5dGVzLmxlbmd0aDsKCQkJCWlmIChvZmZzZXQgJSA4ICE9PSAwKSB7CgkJCQkJb2Zmc2V0ICs9IDggLSAob2Zmc2V0ICUgOCk7CgkJCQl9CgkJCQlyZXR1cm4gcHRyOwoJCQl9OwoKCQkJY29uc3QgYXJnYyA9IHRoaXMuYXJndi5sZW5ndGg7CgoJCQljb25zdCBhcmd2UHRycyA9IFtdOwoJCQl0aGlzLmFyZ3YuZm9yRWFjaCgoYXJnKSA9PiB7CgkJCQlhcmd2UHRycy5wdXNoKHN0clB0cihhcmcpKTsKCQkJfSk7CgkJCWFyZ3ZQdHJzLnB1c2goMCk7CgoJCQljb25zdCBrZXlzID0gT2JqZWN0LmtleXModGhpcy5lbnYpLnNvcnQoKTsKCQkJa2V5cy5mb3JFYWNoKChrZXkpID0+IHsKCQkJCWFyZ3ZQdHJzLnB1c2goc3RyUHRyKGAke2tleX09JHt0aGlzLmVudltrZXldfWApKTsKCQkJfSk7CgkJCWFyZ3ZQdHJzLnB1c2goMCk7CgoJCQljb25zdCBhcmd2ID0gb2Zmc2V0OwoJCQlhcmd2UHRycy5mb3JFYWNoKChwdHIpID0+IHsKCQkJCXRoaXMubWVtLnNldFVpbnQzMihvZmZzZXQsIHB0ciwgdHJ1ZSk7CgkJCQl0aGlzLm1lbS5zZXRVaW50MzIob2Zmc2V0ICsgNCwgMCwgdHJ1ZSk7CgkJCQlvZmZzZXQgKz0gODsKCQkJfSk7CgoJCQkvLyBUaGUgbGlua2VyIGd1YXJhbnRlZXMgZ2xvYmFsIGRhdGEgc3RhcnRzIGZyb20gYXQgbGVhc3Qgd2FzbU1pbkRhdGFBZGRyLgoJCQkvLyBLZWVwIGluIHN5bmMgd2l0aCBjbWQvbGluay9pbnRlcm5hbC9sZC9kYXRhLmdvOndhc21NaW5EYXRhQWRkci4KCQkJY29uc3Qgd2FzbU1pbkRhdGFBZGRyID0gNDA5NiArIDgxOTI7CgkJCWlmIChvZmZzZXQgPj0gd2FzbU1pbkRhdGFBZGRyKSB7CgkJCQl0aHJvdyBuZXcgRXJyb3IoInRvdGFsIGxlbmd0aCBvZiBjb21tYW5kIGxpbmUgYW5kIGVudmlyb25tZW50IHZhcmlhYmxlcyBleGNlZWRzIGxpbWl0Iik7CgkJCX0KCgkJCXRoaXMuX2luc3QuZXhwb3J0cy5ydW4oYXJnYywgYXJndik7CgkJCWlmICh0aGlzLmV4aXRlZCkgewoJCQkJdGhpcy5fcmVzb2x2ZUV4aXRQcm9taXNlKCk7CgkJCX0KCQkJYXdhaXQgdGhpcy5fZXhpdFByb21pc2U7CgkJfQoKCQlfcmVzdW1lKCkgewoJCQlpZiAodGhpcy5leGl0ZWQpIHsKCQkJCXRocm93IG5ldyBFcnJvcigiR28gcHJvZ3JhbSBoYXMgYWxyZWFkeSBleGl0ZWQiKTsKCQkJfQoJCQl0aGlzLl9pbnN0LmV4cG9ydHMucmVzdW1lKCk7CgkJCWlmICh0aGlzLmV4aXRlZCkgewoJCQkJdGhpcy5fcmVzb2x2ZUV4aXRQcm9taXNlKCk7CgkJCX0KCQl9CgoJCV9tYWtlRnVuY1dyYXBwZXIoaWQpIHsKCQkJY29uc3QgZ28gPSB0aGlzOwoJCQlyZXR1cm4gZnVuY3Rpb24gKCkgewoJCQkJY29uc3QgZXZlbnQgPSB7IGlkOiBpZCwgdGhpczogdGhpcywgYXJnczogYXJndW1lbnRzIH07CgkJCQlnby5fcGVuZGluZ0V2ZW50ID0gZXZlbnQ7CgkJCQlnby5fcmVzdW1lKCk7CgkJCQlyZXR1cm4gZXZlbnQucmVzdWx0OwoJCQl9OwoJCX0KCX0KfSkoKTsK"
	FileHTML = "PCFkb2N0eXBlIGh0bWw+CjxodG1sPgo8aGVhZD4KCTxtZXRhIGNoYXJzZXQ9InV0Zi04Ij4KCTx0aXRsZT5HbyB3YXNtPC90aXRsZT4KCTxzY3JpcHQgc3JjPSJ3YXNtX2V4ZWMuanMiPjwvc2NyaXB0Pgo8L2hlYWQ+Cgo8Ym9keT4KCTxzY3JpcHQ+CgkJY29uc3QgZ28gPSBuZXcgR28oKTsKCQlXZWJBc3NlbWJseS5pbnN0YW50aWF0ZVN0cmVhbWluZyhmZXRjaCgibWFpbi53YXNtIiksIGdvLmltcG9ydE9iamVjdCkudGhlbigocmVzdWx0KSA9PiB7CgkJCWdvLnJ1bihyZXN1bHQuaW5zdGFuY2UpOwoJCX0pLmNhdGNoKChlcnIpID0+IHsKCQkJY29uc29sZS5lcnJvcihlcnIpOwoJCX0pOwoJPC9zY3JpcHQ+CjwvYm9keT4KPC9odG1sPgo="
)
