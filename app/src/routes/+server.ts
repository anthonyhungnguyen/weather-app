import { BACKEND_HOST } from '$env/static/private';
import { json } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ url }) {
	const city = url.searchParams.get('city');
	const resp = await fetch(`${BACKEND_HOST}/weather?city=${city}`);
	const data = await resp.json();
	return json(data);
}
