/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/',
        destination: '/ja',
      },
    ];
  }
};

module.exports = nextConfig;
