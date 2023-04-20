describe('Log in.cy.ts', () => {
  it('ensures log in process works correctly', () => {
    cy.visit('http://localhost:4200')

    cy.contains('Log in with Google').click()

    cy.wait(30000)
    cy.contains('Log out').should('exist')
  })
})