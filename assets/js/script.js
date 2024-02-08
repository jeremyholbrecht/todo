// TODO: make it so that each done button is properly associated with it's task
// currently it doesn't matter which button we click it will go from top to buttom instead of it's associated task

const taskDoneButtons = document.querySelectorAll('.taskDoneButton');

taskDoneButtons.forEach(taskDoneButton => {
    taskDoneButton.addEventListener('click', () => {
        const task = document.querySelector('.unfinished');
        task.classList.remove('unfinished');
        task.classList.add('done');
    });
});