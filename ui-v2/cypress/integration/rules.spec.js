describe('/rules', () => {
  it('first load', () => {
    cy.setCookie('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImZhaHJpLmhpZGF5YXRAdGlrZXQuY29tIiwiZXhwIjoxNTg0MDk3NTc3LCJwaWN0dXJlIjoiaHR0cHM6Ly9saDUuZ29vZ2xldXNlcmNvbnRlbnQuY29tLy1FNG5hTlVWMEo4by9BQUFBQUFBQUFBSS9BQUFBQUFBQUFBQS93Y3ZMcXR6RnlpWS9waG90by5qcGcifQ.B2WoqwX6owiYv0d2fbRWn6fIk5MbjvIXlxVNdnvFl0E');
    cy.visit('/rules');
  });
});
