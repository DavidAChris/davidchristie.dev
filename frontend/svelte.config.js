import adapter from '@sveltejs/adapter-static';
// See https://kit.svelte.dev/docs/adapters for more information about adapters.

/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		adapter: adapter({
			pages: '../views',
			assets: '../static',
			fallback: undefined,
			precompress: false,
			strict: true
		})
	}
};

export default config;
