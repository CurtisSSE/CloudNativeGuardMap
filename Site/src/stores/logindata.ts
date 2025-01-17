import { writable } from 'svelte/store';
// import { User } from '../../typedefs/user.js'

// const userdata: User = null

export const loggedInUser = writable('');
// export const userData = readable(userdata)