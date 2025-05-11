import { init } from '@wjfe/n-savant'
import { mount } from 'svelte'
import './app.css'
import App from './Main.svelte'

init();

const app = mount(App, {
  target: document.getElementById('app')!,
})

export default app
