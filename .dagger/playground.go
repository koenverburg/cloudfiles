package main

import "dagger/cloudfiles/internal/dagger"

func (m *Cloudfiles) playgroundContainer() *dagger.Container {
	return dag.Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
		From("ubuntu:22.04").

		WithExec([]string{"sh", "-c",
			"apt-get update && " +
				"apt-get install -y --no-install-recommends " +
				"tmux jq git curl gnupg wget unzip xz-utils " +
				"build-essential libudev-dev pkg-config libssl-dev ca-certificates && " +
				"rm -rf /var/lib/apt/lists/*"}).

		WithExec([]string{"sh", "-c",
			"curl -fsSL https://go.dev/dl/go1.23.5.linux-amd64.tar.gz | tar -xz -C /usr/local"}).

		WithEnvVariable("RUSTUP_HOME", "/usr/local/rustup").
		WithEnvVariable("CARGO_HOME", "/usr/local/cargo").
		WithExec([]string{"sh", "-c",
			"curl -fsSL https://sh.rustup.rs | sh -s -- -y --profile minimal --default-toolchain stable"}).

		WithExec([]string{"sh", "-c",
			"curl -fsSL https://github.com/Schniz/fnm/releases/download/v1.38.1/fnm-linux.zip -o /tmp/fnm.zip && " +
				"unzip -q /tmp/fnm.zip -d /usr/local/bin && " +
				"chmod +x /usr/local/bin/fnm && " +
				"rm /tmp/fnm.zip"}).

		WithExec([]string{"sh", "-c",
			"curl -fsSL https://github.com/zellij-org/zellij/releases/download/v0.41.2/zellij-x86_64-unknown-linux-musl.tar.gz | " +
				"tar -xz -C /usr/local/bin zellij && " +
				"chmod +x /usr/local/bin/zellij"}).

		WithExec([]string{"sh", "-c",
			"curl -fsSL https://nodejs.org/dist/v22.13.1/node-v22.13.1-linux-x64.tar.xz | " +
				"tar -xJ -C /usr/local --strip-components=1"}).

		// ── Environment variables (set once, after all binaries are in place) ─
		WithEnvVariable("GOROOT", "/usr/local/go").
		WithEnvVariable("GOPATH", "/root/go").
		WithEnvVariable("FNM_DIR", "/usr/local/fnm").
		WithEnvVariable("TERM", "xterm-256color").
		WithEnvVariable("PATH",
			"/usr/local/cargo/bin:"+
				"/usr/local/go/bin:"+
				"/root/go/bin:"+
				"/usr/local/bin:"+
				"/usr/bin:/bin:/usr/sbin:/sbin").

		WithExec([]string{"sh", "-c", "cargo install moltis --git https://github.com/moltis-org/moltis"}).


		// WithExec([]string{"sh", "-c",
		// 	"fnm install v25",
		// 	// "fnm use v25",
		// 	// "npm install -g @mariozechner/pi-coding-agent",
		// }).

		WithWorkdir("/root")
}
