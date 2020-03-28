describe('rules_edit', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('rules/get_rules_resp_offset0_limit10_total5.json').as('get_rules_resp_offset0_limit10_total5');
    cy.fixture('rules/get_data_sources_resp.json').as('get_data_sources_resp');
    cy.fixture('rules/get_rules_5_resp.json').as('get_rules_5_resp');

    const resizeObserverLoopErrRe = /^ResizeObserver loop limit exceeded/;
    Cypress.on('uncaught:exception', (err) => {
      if (resizeObserverLoopErrRe.test(err.message)) {
        return false;
      }
    });
  });

  context('delete', () => {
    it('DELETE', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/data_sources', '@get_data_sources_resp')
        .as('get_data_sources');
      cy.route('GET', '/api/rules/5', '@get_rules_5_resp')
        .as('get_rules_5');
      cy.route('DELETE', '/api/rules/5')
        .as('delete_rules_5');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-row-key=5] > .col-action > [data-testid=colgroup-action] > [data-testid=btn-delete]').click();
      cy.get('.ant-popover-buttons > .ant-btn-primary').click({ force: true });

      cy.wait('@delete_rules_5');
    });
  });
});
