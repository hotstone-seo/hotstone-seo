describe('rules_edit', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('rules/get_rules_resp_offset0_limit10_total5.json').as('get_rules_resp_offset0_limit10_total5');
    cy.fixture('rules/get_data_sources_resp.json').as('get_data_sources_resp');
    cy.fixture('rules/post_rules_5_resp.json').as('post_rules_5_resp');
    cy.fixture('rules/get_rules_5_resp.json').as('get_rules_5_resp');
    cy.fixture('rules/put_rules_5_edited_resp.json').as('put_rules_5_edited_resp');

    const resizeObserverLoopErrRe = /^ResizeObserver loop limit exceeded/;
    Cypress.on('uncaught:exception', (err) => {
      if (resizeObserverLoopErrRe.test(err.message)) {
        return false;
      }
    });
  });

  context('edit', () => {
    it('GIVEN Rule without Data Source AND Data Source selected THEN set id_data_source', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/data_sources', '@get_data_sources_resp')
        .as('get_data_sources');
      cy.route('GET', '/api/rules/5', '@get_rules_5_resp')
        .as('get_rules_5');
      cy.route('PUT', '/api/rules', '@put_rules_5_edited_resp')
        .as('put_rules_5');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-row-key=5] > .col-action > [data-testid=colgroup-action] > [data-testid=btn-edit]').click();

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/5`);
      cy.wait('@get_rules_5');

      cy.get('[data-testid=btn-edit]').click();

      cy.get('[data-testid="input-name"]').clear().type('Foo Rule Edited');
      cy.get('[data-testid="input-url-pattern"]').clear().type('foo/rule/edited');
      cy.get('[data-testid="select-data-source-id"]').type('Foo Data Source');
      cy.get('.ant-select-item-option-active').click();

      cy.get('[data-testid="btn-save"]').click();

      cy.wait('@put_rules_5');
      cy.get('@put_rules_5')
        .its('request.body')
        .should('deep.equal', {
          id: 5,
          name: 'Foo Rule Edited',
          url_pattern: '/foo/rule/edited',
          data_source_id: 1,
        });

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/5`);
    });
  });
});
