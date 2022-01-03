import {shallowMount} from '@vue/test-utils'
import ToDoApp from "../../../src/components/ToDoApp.vue"


describe('ToDoApp.vue',function(){

  it('Checking to change the value of input box',async function(){
    const wrapper = shallowMount(ToDoApp);
    const textInput = wrapper.find('#input-box')
    await textInput.setValue('some value')
    expect(wrapper.find('#input-box').element.textContent).toBe("some value")
  })
})

describe('ToDoApp.vue',function(){

  it('Checking when clicked to add button is the input be deleted',async function(){
    const wrapper = shallowMount(ToDoApp);
    const textInput = wrapper.find('#input-box')
    await textInput.setValue('some value')
    const addButton = wrapper.find('#add-button')
    await addButton.trigger('click')
    expect(wrapper.find('#input-box').element.textContent).toBe('')
  })
})

