package service_test

// FIXME:
// func TestClientKeyService_Insert(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	repoMock := repository_mock.NewMockClientKeyRepo(ctrl)
// 	auditSvcMock := service_mock.NewMockAuditTrailService(ctrl)
// 	ctx := context.Background()

// 	svc := service.ClientKeyServiceImpl{
// 		ClientKeyRepo: repoMock,
// 		AuditTrailService: auditSvcMock,
// 	}

// 	newClientKey := repository.ClientKey{Name: "Foo", Prefix: "123", Key: "456"}
// 	repoMock.EXPECT().Insert(ctx, gomock.Any()).Return(newClientKey, nil)

// 	givenClientKey := repository.ClientKey{Name: "Foo"}
// 	data, err := svc.Insert(ctx, givenClientKey)
// 	require.NoError(t, err)
// 	require.Equal(t, "Foo", data.Name)
// 	require.Equal(t, "123", data.Prefix)
// 	require.NotEqual(t, "456" , data.Key)
// 	require.Equal(t, 32, len(data.Key))
// }
