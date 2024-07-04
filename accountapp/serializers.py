from rest_framework import serializers
from accountapp.models import *


class AccountSerializer(serializers.ModelSerializer):
    user = serializers.PrimaryKeyRelatedField(queryset=UserProfile.objects.all(),
                                              required=False)  # ������ ���� user ��������������

    class Meta:
        model = Account
        fields = '__all__'
