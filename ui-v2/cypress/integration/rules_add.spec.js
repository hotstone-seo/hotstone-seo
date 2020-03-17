describe('rules_add', () => {
  beforeEach(() => {
    cy.fixture('cookies.json').then((cookies) => {
      cy.setCookie('token', cookies.token);
    });

    cy.server();

    cy.fixture('rules/get_rules_resp_offset0_limit10_total5.json').as('get_rules_resp_offset0_limit10_total5');
    cy.fixture('rules/get_data_sources_resp.json').as('get_data_sources_resp');
    cy.fixture('rules/post_rules_5_resp.json').as('post_rules_5_resp');
    cy.fixture('rules/get_rules_5_resp.json').as('get_rules_5_resp');
    cy.fixture('rules/post_rules_6_resp.json').as('post_rules_6_resp');
    cy.fixture('rules/get_rules_6_resp.json').as('get_rules_6_resp');

    const resizeObserverLoopErrRe = /^ResizeObserver loop limit exceeded/;
    Cypress.on('uncaught:exception', (err) => {
      if (resizeObserverLoopErrRe.test(err.message)) {
        return false;
      }
    });
  });

  context('add', () => {
    it('GIVEN error validation THEN no post request', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/data_sources', '@get_data_sources_resp')
        .as('get_data_sources');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-testid="btn-new-rule"]').click();
      cy.wait('@get_data_sources');

      cy.get('[data-testid="btn-save"]').click();

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/new`);
    });

    it('GIVEN no datasource selected THEN no id_data_source', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/data_sources', '@get_data_sources_resp')
        .as('get_data_sources');
      cy.route('POST', '/api/rules', '@post_rules_5_resp')
        .as('post_rules');
      cy.route('GET', '/api/rules/5', '@get_rules_5_resp')
        .as('get_rules_5');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-testid="btn-new-rule"]').click();
      cy.wait('@get_data_sources');

      cy.get('[data-testid="input-name"]').type('Foo Rule');
      cy.get('[data-testid="input-url-pattern"]').type('/foo/rule');

      cy.get('[data-testid="btn-save"]').click();

      cy.get('@post_rules')
        .its('request.body')
        .should('deep.equal', {
          name: 'Foo Rule',
          url_pattern: '/foo/rule',
        });

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/5`);
      cy.wait('@get_rules_5');
    });

    it('GIVEN datasource selected THEN set id_data_source', () => {
      cy.route('GET', '/api/rules?_offset=0&_limit=10', '@get_rules_resp_offset0_limit10_total5')
        .as('get_rules_offset0_limit10');
      cy.route('GET', '/api/data_sources', '@get_data_sources_resp')
        .as('get_data_sources');
      cy.route('POST', '/api/rules', '@post_rules_6_resp')
        .as('post_rules');
      cy.route('GET', '/api/rules/6', '@get_rules_6_resp')
        .as('get_rules_6');

      cy.visit('/rules');
      cy.wait('@get_rules_offset0_limit10');

      cy.get('[data-testid="btn-new-rule"]').click();
      cy.wait('@get_data_sources');

      cy.get('[data-testid="input-name"]').type('Dummy Rule');
      cy.get('[data-testid="input-url-pattern"]').type('/dummy/rule');
      cy.get('[data-testid="select-data-source-id"]').type('Foo Data Source');

      cy.get('.ant-select-item-option-active').click();
      cy.get('[data-testid="btn-save"]').click();

      cy.get('@post_rules')
        .its('request.body')
        .should('deep.equal', {
          name: 'Dummy Rule',
          url_pattern: '/dummy/rule',
          data_source_id: 1,
        });

      cy.url().should('eq', `${Cypress.config().baseUrl}/rules/6`);
      cy.wait('@get_rules_6');
    });
  });
});
