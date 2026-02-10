// TODO TASK 8: Write the mutation for ToggleTaskComplete
// Should toggle a task's completion status and return the updated task
export const TOGGLE_TASK_COMPLETE = gql`
  mutation ToggleTaskComplete($id: ID!) {
    toggleTaskComplete(id: $id) {
      id
      title
      description
      completed
      priority
      createdAt
      updatedAt
    }
  }
`