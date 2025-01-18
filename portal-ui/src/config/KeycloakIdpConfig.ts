import { KeycloakIdpConfig, KeycloakIdpHint } from '../types/KeycloakIdpConfig';

const keycloakIdpConfig: KeycloakIdpConfig = {
  clientId: 'idp-1',
  clientSecret: 'c5CH107aKaxJMaJ5WvdW7rlYjV4IaFCd',
  issuer: 'http://localhost:8080/realms/idp-1',
  redirectUri: 'http://localhost:3000/callback',
  scopes: ['openid', 'profile', 'email'],
  postLogoutRedirectUri: 'https://localhost:3000/logout-callback'
};

const keycloakIdpHints: KeycloakIdpHint[] = [
  {
    providerId: 'keycloak-oidc',
    displayName: 'KEYCLOAK'
  }
];

export {
  keycloakIdpHints,
  keycloakIdpConfig
 };
