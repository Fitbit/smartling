class Smartling < Formula
  desc "CLI to upload and download translations"
  homepage "https://github.com/mdreizin/smartling"
  url "https://github.com/mdreizin/smartling/archive/v0.4.3.tar.gz"
  sha256 "ac246a5d2d31925a31b9ed8e2cb195edd97faf3a2e43e5726ddb2dffe61eabaa"

  head "https://github.com/mdreizin/smartling.git"

  depends_on "go"
  depends_on "glide"

  def install
    ENV["GOPATH"] = buildpath
    glidepath = buildpath/".glide"
    smartlingpath = buildpath/"src/github.com/mdreizin/smartling"
    smartlingpath.install buildpath.children

    cd smartlingpath do
      system "glide", "--home", "#{glidepath}", "install", "--skip-test"
      system "go", "build", "-o", "smartling", "./cli/..."
      bin.install "smartling"
      prefix.install_metafiles
    end
  end

  test do
    version = pipe_output("#{bin}/smartling --version")
    assert_match version.to_s, version
  end
end
