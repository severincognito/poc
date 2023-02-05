import adapter from '@sveltejs/adapter-static';

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: '../static',
			assets: '../static',
      		fallback: null,
      		precompress: false,
      		strict: true
		})
	}
};

export default config;