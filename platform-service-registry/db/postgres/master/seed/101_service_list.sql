do language plpgsql
$$
    begin
        insert into
            platform_service_registry.service (name, namespace, created_at)
        values
            ('platform-service-registry', 'platform', now()),
            ('platform-service-asset-registry', 'platform', now()),
            ('platform-service-generator', 'platform', now()),
            ('platform-service-desktop', 'platform', now()),
            ('platform-service-deployment-spec-generator', 'platform', now()),

            ('postgres-control-plane', 'dbaas', now()),
            ('mongodb-control-plane', 'dbaas', now()),
            ('redis-control-plane', 'dbaas', now()),

            ('kubernetes-control-plane', 'compute', now()),
            ('kvm-control-plane', 'compute', now()),

            ('network-control-plane', 'networking', now());
    end;
$$;
