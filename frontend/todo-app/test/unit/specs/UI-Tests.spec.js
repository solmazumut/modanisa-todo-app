import {shallowMount} from '@vue/test-utils'
import ToDoApp from "../../../src/components/ToDoApp.vue"

describe('ToDoApp.vue',function(){
    it('Checking UI - #input-box',function(){
        const wrapper = shallowMount(ToDoApp);
        const inputBox = wrapper.find('#input-box');
        expect(inputBox.exists()).toBe(true)
    })
})

describe('ToDoApp.vue',function(){
  it('Checking UI - #add-button',function(){
    const wrapper = shallowMount(ToDoApp);
    const addButton = wrapper.find('#add-button');
    expect(addButton.exists()).toBe(true)
  })
})

describe('ToDoApp.vue',function(){
  it('Checking UI - #to-do-table',function(){
    const wrapper = shallowMount(ToDoApp);
    const toDoTable = wrapper.find('#to-do-table');
    expect(toDoTable.exists()).toBe(true)
  })
})


describe('ToDoApp.vue',function(){
  it('Checking UI - #reset-button',function(){
    const wrapper = shallowMount(ToDoApp);
    const resetButton = wrapper.find('#reset-button');
    expect(resetButton.exists()).toBe(true)
  })
})
