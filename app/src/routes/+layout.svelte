<script>
	import { page } from '$app/stores';
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
	import { ChartPieSolid } from 'flowbite-svelte-icons';
	import { Footer, FooterCopyright, FooterLinkGroup, FooterLink } from 'flowbite-svelte';

	import '../app.css';
	$: activeUrl = $page.url.pathname;
</script>

<div class="flex flex-row h-screen">
	<div class="basis-1/6 bg-black mr-4">
		<div class="text-white font-bold text-center py-4 border-b border-white">The Weather App</div>
		<Sidebar>
			<SidebarWrapper>
				<SidebarGroup>
					<SidebarItem href="/weather" active={activeUrl === '/weather'} label="Weather"
						><svelte:fragment slot="icon">
							<ChartPieSolid class="w-5 h-5" />
						</svelte:fragment></SidebarItem
					>
				</SidebarGroup>
			</SidebarWrapper>
		</Sidebar>
	</div>
	<div class="basis-5/6">
		<slot />
	</div>
</div>
<Footer class="absolute bottom-0 left-0 z-20 w-full">
	<FooterCopyright href="/" by="HungNP" year={2023} />
	<FooterLinkGroup
		ulClass="flex flex-wrap items-center mt-3 text-sm text-gray-500 dark:text-gray-400 sm:mt-0"
	>
		<FooterLink href="/">About</FooterLink>
		<FooterLink href="/">Privacy Policy</FooterLink>
		<FooterLink href="/">Licensing</FooterLink>
		<FooterLink href="/">Contact</FooterLink>
	</FooterLinkGroup>
</Footer>
