import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],

	theme: {
		screens: {
			msm: "375px",
			mbl: "430px",
			sm: "640px",
			md: "768px",
			lg: "1024px",
			xl: "1280px",
			"2xl": "1536px"
		},
		// colors: {
		// },
		extend: {
			colors: {
				"base2": "#050609",
				// "crust": "#0d1321",
				"crust": "hsl(240deg 8% 8%)",
				"navbar": "hsl(240deg 8% 16%)",
				"mantle": "hsl(240deg 8% 9%)",
				"base": "hsl(210, 11%, 11%)",
				// "surface0": "#343a40",
				"surface0": "hsl(240deg 8% 18%)",
				"surface01": "hsl(240deg 8% 24%)",
				"surface1": "hsl(240deg 8% 24%)",
				"surface2": "hsl(240deg 8% 32%)",
				"selected-order": "hsl(287 17% 30%)",
				"selected-order-checkbox": "hsl(287 64% 70%)",
				"selected-order-checkbox-hover": "hsl(287 64% 74%)",
				"selected-order-text": "hsl(287 64% 80%)",
				"text": "hsl(60, 10%, 92%)",
				"subtext2": "hsl(240, 10%, 88%)",
				"subtext1": "hsl(240, 10%, 78%)",
				"subtext0": "hsl(240, 10%, 64%)",
				"success": "#C0D684",
				"success2": "#CBEAA6",
				"mauve": "#BAA5FF",
				"sblue": "#8ecae6",
				"sgreen": "#80ed99",
			}
		}
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
