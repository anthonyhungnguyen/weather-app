<script lang="ts">
	export let onClick: (event: MouseEvent) => void;
	export let type = 'primary';

	let classes = '';
	if (type === 'primary') {
		classes += 'bg-blue-500 text-white';
	} else if (type === 'secondary') {
		classes += 'bg-gray-500 text-white';
	}
</script>

<button class="{classes} font-bold py-2 px-4 rounded" on:click|preventDefault={onClick}>
	<slot />
</button>
