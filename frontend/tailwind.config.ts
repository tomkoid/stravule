import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		screens: {
			msm: "375px",
			sm: "640px",
			md: "768px",
			lg: "1024px",
			xl: "1280px",
			"2xl": "1536px"
		},
		extend: {}
	},
	future: {
		hoverOnlyWhenSupported: true,
	},
	plugins: [
		require("@catppuccin/tailwindcss")({
			// prefix to use, e.g. `text-pink` becomes `text-ctp-pink`.
			// default is `false`, which means no prefix
			prefix: "ctp",
			// which flavour of colours to use by default, in the `:root`
			defaultFlavour: "mocha",
		}),
	]
} satisfies Config;
